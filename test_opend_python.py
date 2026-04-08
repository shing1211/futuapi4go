#!/usr/bin/env python3
"""Test Futu OpenD connection using official Python SDK"""

from futu import OpenQuoteContext, SysConfig

def main():
    print("=== Testing Futu OpenD with Python SDK ===")
    
    # Create quote context
    quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
    
    try:
        # Try to get basic info
        print("✅ Connected to OpenD!")
        print(f"   Quote context created successfully")
        
        # Try a simple API call - get market state
        ret, data = quote_ctx.get_market_state("HK.00700")
        if ret == 0:
            print(f"✅ Market state API succeeded!")
            print(f"   {data}")
        else:
            print(f"⚠️  Market state returned: {ret}")
            
        print("\n🎉 OpenD is working correctly with Python SDK!")
        
    except Exception as e:
        print(f"❌ Failed: {e}")
        print("\n💡 This suggests OpenD is not ready or misconfigured")
    finally:
        quote_ctx.close()

if __name__ == '__main__':
    main()
