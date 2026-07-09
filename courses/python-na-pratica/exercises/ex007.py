numbers = []
sum = 0.0

for i in range(5):
    numbers.append(float(input("Digite o numero: ")))
    sum += numbers[i]

print(f"A media dos numeros digitados foi {sum / 5}")
