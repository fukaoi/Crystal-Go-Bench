require "openssl"

def main
  hash = OpenSSL::Digest.new("SHA256")
  0.upto(1000000) do |i| 
    hash << (Time.now.millisecond + i).to_s
    puts hash.hexdigest  
  end
end

main