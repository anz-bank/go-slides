from subprocess import check_output
import re
import matplotlib.pyplot as plt
import csv
import pandas
import os
import time
filename = "responsetime.csv"
address = "localhost"
request_range = [1000, 20000, 1000]

def testport(address, port, numRequests):
    timeout = 2000
    command = f"ab -n {numRequests} -c {numRequests} -s {timeout} -q -r -e temp.csv {address}:{port}/ping"
    print(command)
    os.system(command)
    file = open("temp.csv", "r")
    os.system("rm temp.csv")
    return file

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
        f = testport(address, 9090, numRequests)
        go_list = read_file(f)
        f = testport(address, 8080, numRequests)
        java_list = read_file(f)
        c.writerow([numRequests] + go_list + java_list)
    file.close()

def generate_graph(filename, request_range):
    data = pandas.read_csv(filename)
    # plt.figure(figsize=(10,6))
    plt.plot(data.numRequests.tolist(), data.java_50.div(1000).tolist(), "o", label = "Java Median", color='#ff9721')
    plt.plot(data.numRequests.tolist(), data.java_95.div(1000).tolist(), "^", label = "Java 95th percentile", color='#ff9721')
    plt.plot(data.numRequests.tolist(), data.go_50.div(1000).tolist(), "o", label = "Go Median", color = '#4f92ff')
    plt.plot(data.numRequests.tolist(), data.go_95.div(1000).tolist(), "^", label = "Go 95th percentile", color = '#4f92ff')
    bottom, top = plt.ylim()
    # plt.ylim(top = top + 20)
    plt.title('Java vs Go server response time')
    plt.xlabel('Number of concurrent requests')
    plt.ylabel('Time taken (s)')
    plt.legend()
    plt.xticks(ticks = [x for x in range(request_range[0], request_range[1], request_range[2]*2)])
    plt.savefig(os.path.abspath('..')+"/content/_img/JavaVsGoLoadTest.png", dpi=300, bbox_inches='tight', figsize=(50,25))


if os.path.isfile(filename) is False:
    compare(filename, address, request_range)
generate_graph(filename, request_range)
