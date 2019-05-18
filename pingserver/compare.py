from subprocess import check_output
import re
import matplotlib.pyplot as plt
import csv
import pandas

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


file = open("Responsetime.csv", "w+")
c = csv.writer(file)
request_range = range(1000, 10000, 1000)
header = ["number_of_requests", "Go_median_time", "Go_max_error", "Go_min_error", "Java_median_time", "Java_max_error", "Java_min_error"]
address = "192.168.0.19"


for num in request_range:
    go_median, go_max, go_min = testport(address, 9090, num)
    go_min_error = go_median - go_min
    go_max_error = go_max - go_median
    java_median, java_max, java_min = testport(address, 8080, num)
    java_min_error = java_median - java_min
    java_max_error = java_max - java_median
    c.writerow([num, go_median, go_max_error, go_min_error, java_median, java_max_error, java_min_error])
file.close()

# Plotting
data = pandas.read_csv("Responsetime.csv", names=header)
plt.errorbar(data.number_of_requests.tolist(), data.Go_median_time.tolist(), yerr = [data.Go_min_error.tolist(), data.Go_max_error.tolist()], fmt = "o", label = "Go server")
plt.errorbar(data.number_of_requests.tolist(), data.Java_median_time.tolist(), yerr = [data.Java_min_error.tolist(), data.Java_max_error.tolist()], fmt = "o", label = "Java server")

plt.title('Java vs Go server response time')
plt.xlabel('Number of concurrent requests')
plt.ylabel('Time taken (s)')
plt.legend()
plt.savefig("JavaVsGoLoadTest")
plt.show()
