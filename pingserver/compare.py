from subprocess import check_output
import re
import matplotlib.pyplot as plt
import csv
import pandas
import os
import time
filename = "responsetime.csv"
address = "localhost"
request_range = [1000, 21000, 1000]

def testport(address, port, numRequests, filename_full):
    timeout = 2000
    command = f"ab -n {numRequests} -c {numRequests} -s {timeout} -q -r -e {filename_full} {address}:{port}/ping"
    print(command)
    check_output(command, shell = True)
    return open(filename_full, "r")

def read_file(file):
    reader = pandas.read_csv(file)
    time_list = [reader.loc[i]["Time in ms"] for i in range(0, 101, 1)]
    return time_list

def compare(filename, address, request_range):
    file = open(filename, "w+")
    c = csv.writer(file)
    header = ["numRequests"] + ["go_"+ str(i) for i in range(0, 101, 1)] + ["java_"+ str(i) for i in range(0, 101, 1)]
    c.writerow (header)
    for numRequests in range(request_range[0], request_range[1], request_range[2]):
        filename_full = f"responses/{filename}_9090_{numRequests}.csv"
        f = testport(address, 9090, numRequests, filename_full)
        go_list = read_file(f)

        filename_full = f"responses/{filename}_8080_{numRequests}.csv"
        f = testport(address, 8080, numRequests, filename_full)
        java_list = read_file(f)
        c.writerow([numRequests] + go_list + java_list)
    file.close()

def generate_graph(filename, request_range):
    data = pandas.read_csv(filename)
    plt.figure(figsize=(10,6))
    plt.plot(data.numRequests.tolist(), data.java_50.tolist(), "o", label = "Java Median", color='#ff9721')
    plt.plot(data.numRequests.tolist(), data.java_90.tolist(), "^", label = "Java 90th percentile", color='#ff9721')
    plt.plot(data.numRequests.tolist(), data.go_50.tolist(), "o", label = "Go Median", color = '#4f92ff')
    plt.plot(data.numRequests.tolist(), data.go_90.tolist(), "^", label = "Go 90th percentile", color = '#4f92ff')
    bottom, top = plt.ylim()
    plt.ylim(top = top + 20)
    plt.title('Java vs Go server response time')
    plt.xlabel('Number of concurrent requests')
    plt.ylabel('Time taken (s)')
    plt.legend()
    plt.xticks(ticks = [x for x in range(request_range[0], request_range[1], request_range[2]*2)])
    plt.savefig(os.path.abspath('..')+"/content/_img/JavaVsGoLoadTest.png", dpi=300, bbox_inches='tight', figsize=(50,25))


if os.path.isfile(filename) is False:
    try:
        check_output("netstat -an | grep 9090 | grep -i listen", shell = True)
    except:
        os.system("go run go/main.go > /dev/null &")
    try:
        check_output("netstat -an | grep 8080 | grep -i listen", shell = True)
    except:
        os.system("cd java && ./gradlew clean bootRun > /dev/null &")
        time.sleep(10) # wait for the server to start
    compare(filename, address, request_range)
generate_graph(filename, request_range)
