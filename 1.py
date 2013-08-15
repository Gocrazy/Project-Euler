def sumOfMultiples(factor):
	sum = 0
	for i in range(1,1000):
		if i % factor == 0:
			sum += i
	return sum

print sumOfMultiples(5) + sumOfMultiples(3) - sumOfMultiples(15)

