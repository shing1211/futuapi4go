# Very simple raw socket test - send minimal data and see if OpenD responds
import socket
import struct
import time

def main():
    print("=== Raw TCP Test for Futu OpenD ===")
    
    # Connect
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.settimeout(5.0)
    
    try:
        sock.connect(('127.0.0.1', 11111))
        print("✅ TCP connected")
    except Exception as e:
        print(f"❌ TCP connect failed: {e}")
        return
    
    # Build minimal InitConnect packet
    # ProtoID: 1001 (little endian)
    # ProtoFmt: 1 (Protobuf)
    # ProtoVer: 1
    # SerialNo: 1
    # BodyLen: 2 (minimal body)
    # Body: {clientVer: 1, clientID: ""}
    
    header = struct.pack('<2sIII I 20s 8s',
        b'FT',                    # Magic (2 bytes)
        1001,                     # ProtoID (4 bytes)
        1,                        # ProtoFmt (4 bytes) - Protobuf
        1,                        # ProtoVer (2 bytes) + padding (2 bytes)
        1,                        # SerialNo (4 bytes)
        2,                        # BodyLen (4 bytes)
        b'\x00' * 20,            # SHA1 (20 bytes)
        b'\x00' * 8              # Reserved (8 bytes)
    )
    
    # Minimal protobuf: {clientVer: 1, clientID: ""}
    body = bytes([0x08, 0x01, 0x12, 0x00])
    
    print(f"📤 Sending {len(header)} byte header + {len(body)} byte body")
    sock.sendall(header + body)
    
    print("📥 Waiting for response...")
    try:
        # Try to read response
        data = sock.recv(1024)
        print(f"✅ Received {len(data)} bytes")
        print(f"   Data: {data.hex()}")
        
        if len(data) >= 48:
            # Parse header
            magic = data[0:2]
            proto_id = struct.unpack('<I', data[2:6])[0]
            body_len = struct.unpack('<I', data[13:17])[0]
            
            print(f"   Magic: {magic}")
            print(f"   ProtoID: {proto_id}")
            print(f"   BodyLen: {body_len}")
            
    except socket.timeout:
        print("❌ Timeout - no response from OpenD")
        print("\n💡 Possible reasons:")
        print("   1. OpenD not logged in to Futu account")
        print("   2. OpenD in 'trading lock' state")
        print("   3. API access disabled in settings")
        print("   4. Firewall/antivirus blocking")
    except Exception as e:
        print(f"❌ Error: {e}")
    finally:
        sock.close()
    
    print("\n=== Test Complete ===")

if __name__ == '__main__':
    main()
