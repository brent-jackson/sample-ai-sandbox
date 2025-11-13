# Enhanced Python FastMCP Server
from fastmcp import FastMCP

mcp = FastMCP("Demo 🚀")

@mcp.tool
def add(a: int, b: int) -> int:
    """Add two numbers together"""
    return a + b

@mcp.tool
def subtract(a: int, b: int) -> int:
    """Subtract second number from first number"""
    return a - b

@mcp.tool
def multiply(a: int, b: int) -> int:
    """Multiply two numbers together"""
    return a * b

@mcp.tool
def divide(a: float, b: float) -> float:
    """Divide first number by second number"""
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b

@mcp.tool
def greet(name: str, greeting: str = "Hello") -> str:
    """Greet someone with a customizable greeting"""
    return f"{greeting}, {name}!"

@mcp.tool
def calculate_factorial(n: int) -> int:
    """Calculate the factorial of a positive integer"""
    if n < 0:
        raise ValueError("Factorial is not defined for negative numbers")
    if n == 0 or n == 1:
        return 1
    
    result = 1
    for i in range(2, n + 1):
        result *= i
    return result

if __name__ == "__main__":
    mcp.run()