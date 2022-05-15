redis-benchmark -d 10 -t get,set
====== SET ======
10  100000 requests completed in 1.36 seconds	  73637.70 requests per second
20  100000 requests completed in 1.31 seconds	  76511.09 requests per second
50  100000 requests completed in 1.41 seconds     71022.73 requests per second
100 100000 requests completed in 1.26 seconds     79113.92 requests per second
200 100000 requests completed in 1.29 seconds     77700.08 requests per second
1k  100000 requests completed in 1.32 seconds     75585.79 requests per second
5k  100000 requests completed in 1.34 seconds     74515.65 requests per second


====== GET ======
10  100000 requests completed in 1.30 seconds	  76745.97 requests per second
20  100000 requests completed in 1.30 seconds     76982.29 requests per second
50  100000 requests completed in 1.40 seconds     71428.57 requests per second
100 100000 requests completed in 1.27 seconds     78740.16 requests per second
200 100000 requests completed in 1.27 seconds     78554.59 requests per second
1k  100000 requests completed in 1.31 seconds     76511.09 requests per second
5k  100000 requests completed in 1.39 seconds     72046.11 requests per second





相同长度的value在写入数量越多情况下，平均每个value占用内存多