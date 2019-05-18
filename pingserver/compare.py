from subprocess import check_output
import re
import matplotlib.pyplot as plt
import csv
import pandas
import os
import time
def median_time_regex(input):
    input = str(input)
    regex = "(?:Total:)([\d.\s]*)"
    input = re.search(regex, input).group(1)
    nums = [float(num) for num in input.split()]
    return nums[3]/1000, nums[4]/1000, nums[0]/1000 # this returns: median number, max time, min time

def testport(address, port, numRequests):
    timeout = 2000
    total_requests = "Total:[^\\n]*"
    command = f"ab -n {numRequests} -c {numRequests} -s {timeout} -q -r {address}:{port}/ping | grep -e \"{total_requests}\""
    print(command)
    time_taken = check_output(command, shell = True)
    median_time, max_time, min_time  = median_time_regex(time_taken)
    return median_time, max_time, min_time

def compare(filename, address):
    file = open(filename, "w+")
    c = csv.writer(file)
    request_range = range(1000, 20000, 1000)
    for num in request_range:
        go_median, go_max, go_min = testport(address, 9090, num)
        go_min_error = go_median - go_min
        go_max_error = go_max - go_median
        java_median, java_max, java_min = testport(address, 8080, num)
        java_min_error = java_median - java_min
        java_max_error = java_max - java_median
        c.writerow([num, go_median, go_max_error, go_min_error, java_median, java_max_error, java_min_error])
    file.close()

def generate_graph(filename):
    header = ["number_of_requests", "Go_median_time", "Go_max_error", "Go_min_error", "Java_median_time", "Java_max_error", "Java_min_error"]
    data = pandas.read_csv(filename, names=header)
    plt.figure(figsize=(10,6))
    plt.errorbar(data.number_of_requests.tolist(), data.Go_median_time.tolist(), yerr = [data.Go_min_error.tolist(), data.Go_max_error.tolist()], fmt = "o", label = "Go server")
    plt.errorbar(data.number_of_requests.tolist(), data.Java_median_time.tolist(), yerr = [data.Java_min_error.tolist(), data.Java_max_error.tolist()], fmt = "o", label = "Java server")
    bottom, top = plt.ylim()
    plt.ylim(top = top + 20)
    plt.title('Java vs Go server response time')
    plt.xlabel('Number of concurrent requests')
    plt.ylabel('Time taken (s)')
    plt.legend()
    plt.xticks(ticks = [x for x in range(1000, 21000, 2000)],rotation=0)
    plt.savefig(os.path.abspath('..')+"/content/_img/JavaVsGoLoadTest.png", dpi=300, bbox_inches='tight', figsize=(50,25))
filename = "responsetime.csv"
address = "localhost"
if os.path.isfile(filename) is False:
    try:
        go_status = check_output("curl localhost:9090/ping", shell= True)
    except:
        os.system("go run go/main.go &")
    try:
        java_status = check_output("curl localhost:8080/ping", shell= True)
    except:
        # os.system("cd java")
        os.system("cd java && ./gradlew clean bootRun &")
        time.sleep(10) # wait for the server to start
    compare(filename, address)
    # If Responsetime doesn't exist, create is, otherwise just use what's already there


generate_graph(filename)
