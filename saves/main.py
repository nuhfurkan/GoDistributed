import sys

# The first element in sys.argv is the script name itself
script_name = sys.argv[0]

# The rest of the elements are the command-line parameters
parameters = sys.argv[1:]

sum = 0

for elem in parameters:
    sum += int(elem)

if sum != 0:
    float_number = 1.0/sum
else:
    float_number = 0.0

# Write the floating-point number to sys.stdout
sys.stdout.write(str(float_number))