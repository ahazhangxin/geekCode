```====== SET ======
  100000 requests completed in 1.18 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.071 milliseconds (cumulative count 1)
50.000% <= 0.295 milliseconds (cumulative count 52086)
75.000% <= 0.391 milliseconds (cumulative count 76018)
87.500% <= 0.479 milliseconds (cumulative count 87506)
93.750% <= 0.567 milliseconds (cumulative count 93758)
96.875% <= 0.655 milliseconds (cumulative count 96940)
98.438% <= 0.759 milliseconds (cumulative count 98507)
99.219% <= 0.855 milliseconds (cumulative count 99253)
99.609% <= 0.951 milliseconds (cumulative count 99623)
99.805% <= 1.031 milliseconds (cumulative count 99809)
99.902% <= 1.127 milliseconds (cumulative count 99904)
99.951% <= 51.135 milliseconds (cumulative count 99954)
99.976% <= 51.487 milliseconds (cumulative count 99976)
99.988% <= 51.903 milliseconds (cumulative count 99990)
99.994% <= 51.967 milliseconds (cumulative count 99994)
99.997% <= 52.031 milliseconds (cumulative count 99997)
99.998% <= 52.063 milliseconds (cumulative count 99999)
99.999% <= 52.095 milliseconds (cumulative count 100000)
100.000% <= 52.095 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.052% <= 0.103 milliseconds (cumulative count 52)
16.138% <= 0.207 milliseconds (cumulative count 16138)
54.644% <= 0.303 milliseconds (cumulative count 54644)
78.627% <= 0.407 milliseconds (cumulative count 78627)
89.632% <= 0.503 milliseconds (cumulative count 89632)
95.472% <= 0.607 milliseconds (cumulative count 95472)
97.825% <= 0.703 milliseconds (cumulative count 97825)
98.930% <= 0.807 milliseconds (cumulative count 98930)
99.478% <= 0.903 milliseconds (cumulative count 99478)
99.766% <= 1.007 milliseconds (cumulative count 99766)
99.893% <= 1.103 milliseconds (cumulative count 99893)
99.927% <= 1.207 milliseconds (cumulative count 99927)
99.943% <= 1.303 milliseconds (cumulative count 99943)
99.947% <= 1.407 milliseconds (cumulative count 99947)
99.949% <= 1.503 milliseconds (cumulative count 99949)
99.950% <= 1.607 milliseconds (cumulative count 99950)
99.951% <= 51.103 milliseconds (cumulative count 99951)
100.000% <= 52.127 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 84459.46 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.351     0.064     0.295     0.599     0.823    52.095
====== GET ======
  100000 requests completed in 1.12 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.071 milliseconds (cumulative count 1)
50.000% <= 0.287 milliseconds (cumulative count 52226)
75.000% <= 0.367 milliseconds (cumulative count 75475)
87.500% <= 0.447 milliseconds (cumulative count 88008)
93.750% <= 0.519 milliseconds (cumulative count 93916)
96.875% <= 0.591 milliseconds (cumulative count 96965)
98.438% <= 0.671 milliseconds (cumulative count 98469)
99.219% <= 0.767 milliseconds (cumulative count 99253)
99.609% <= 0.911 milliseconds (cumulative count 99613)
99.805% <= 1.151 milliseconds (cumulative count 99806)
99.902% <= 1.999 milliseconds (cumulative count 99903)
99.951% <= 18.687 milliseconds (cumulative count 99952)
99.976% <= 19.119 milliseconds (cumulative count 99976)
99.988% <= 19.551 milliseconds (cumulative count 99988)
99.994% <= 19.663 milliseconds (cumulative count 99995)
99.997% <= 19.695 milliseconds (cumulative count 99997)
99.998% <= 19.823 milliseconds (cumulative count 99999)
99.999% <= 19.839 milliseconds (cumulative count 100000)
100.000% <= 19.839 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.047% <= 0.103 milliseconds (cumulative count 47)
14.283% <= 0.207 milliseconds (cumulative count 14283)
58.185% <= 0.303 milliseconds (cumulative count 58185)
82.673% <= 0.407 milliseconds (cumulative count 82673)
92.886% <= 0.503 milliseconds (cumulative count 92886)
97.389% <= 0.607 milliseconds (cumulative count 97389)
98.797% <= 0.703 milliseconds (cumulative count 98797)
99.404% <= 0.807 milliseconds (cumulative count 99404)
99.605% <= 0.903 milliseconds (cumulative count 99605)
99.695% <= 1.007 milliseconds (cumulative count 99695)
99.781% <= 1.103 milliseconds (cumulative count 99781)
99.820% <= 1.207 milliseconds (cumulative count 99820)
99.842% <= 1.303 milliseconds (cumulative count 99842)
99.861% <= 1.407 milliseconds (cumulative count 99861)
99.875% <= 1.503 milliseconds (cumulative count 99875)
99.881% <= 1.607 milliseconds (cumulative count 99881)
99.907% <= 2.007 milliseconds (cumulative count 99907)
99.925% <= 2.103 milliseconds (cumulative count 99925)
99.950% <= 3.103 milliseconds (cumulative count 99950)
99.975% <= 19.103 milliseconds (cumulative count 99975)
100.000% <= 20.111 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 89686.10 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.323     0.064     0.287     0.543     0.727    19.839
```