-- simple_calculator.lua
function add(a, b) return a + b end

function subtract(a, b) return a - b end

function multiply(a, b) return a * b end

function divide(a, b) if b ~= 0 then return a / b else return "Error: Division by zero" end end

-- Test the simple calculator functions print("Addition: ", add(10, 5))
print("Subtraction: ", subtract(10, 5)) print("Multiplication: ",
multiply(10, 5)) print("Division: ", divide(10, 5)) print("Division by zero:
", divide(10, 0))


