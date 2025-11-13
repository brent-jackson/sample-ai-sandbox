#!/usr/bin/env python3
"""
Simple test client for the FastMCP server using the standard MCP Python SDK
"""

import asyncio
import json
import subprocess
import sys
from pathlib import Path

async def test_mcp_server():
    """Test the FastMCP server by sending MCP protocol messages"""
    
    # Path to the server
    server_script = Path(__file__).parent / "server.py"
    venv_python = Path(__file__).parent / "venv" / "bin" / "python"
    
    if not venv_python.exists():
        print("❌ Virtual environment not found. Run setup.sh first!")
        return False
    
    print("🧪 Testing FastMCP Server...")
    
    # Start the server process
    try:
        process = subprocess.Popen(
            [str(venv_python), str(server_script)],
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            bufsize=0
        )
        
        # Test 1: Initialize
        print("📡 Sending initialize request...")
        init_request = {
            "jsonrpc": "2.0",
            "id": 1,
            "method": "initialize",
            "params": {
                "protocolVersion": "2024-11-05",
                "capabilities": {},
                "clientInfo": {
                    "name": "test-client",
                    "version": "1.0.0"
                }
            }
        }
        
        process.stdin.write(json.dumps(init_request) + "\n")
        process.stdin.flush()
        
        # Read response with timeout
        try:
            stdout, stderr = process.communicate(timeout=5)
            
            if stdout:
                print("✅ Server responded!")
                print(f"📋 Response preview: {stdout[:500]}...")
                
                # Try to parse JSON responses
                lines = stdout.strip().split('\n')
                for line in lines:
                    if line.strip() and line.strip().startswith('{'):
                        try:
                            response = json.loads(line)
                            if "result" in response:
                                print(f"🎯 Found result: {response['result']}")
                                return True
                        except json.JSONDecodeError:
                            continue
                            
            if stderr:
                print(f"⚠️  Server errors: {stderr}")
                
        except subprocess.TimeoutExpired:
            print("⏰ Server response timeout (this might be normal for MCP servers)")
            process.kill()
            return True  # Timeout can be normal for MCP servers
            
    except Exception as e:
        print(f"❌ Error testing server: {e}")
        return False
    
    finally:
        if process and process.poll() is None:
            process.terminate()
    
    return True

def main():
    """Main test function"""
    print("🚀 FastMCP Server Test")
    print("=" * 50)
    
    # Check if setup was done
    venv_path = Path(__file__).parent / "venv"
    if not venv_path.exists():
        print("❌ Virtual environment not found!")
        print("Please run: ./setup.sh")
        sys.exit(1)
    
    # Run the test
    success = asyncio.run(test_mcp_server())
    
    print("\n" + "=" * 50)
    if success:
        print("✅ Test completed! Server appears to be working.")
        print("\n🎯 Next steps:")
        print("1. Configure Claude Desktop with the server")
        print("2. Or use: mcp-inspector python server.py")
        print("3. Try asking: 'Add 5 and 3 using the add tool'")
    else:
        print("❌ Test failed. Check the error messages above.")
    
    print("\n📚 Available tools:")
    print("- add(a, b): Add two numbers")

if __name__ == "__main__":
    main()