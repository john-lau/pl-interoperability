#!/usr/bin/env ruby

def mem(pid); `ps p #{pid} -o rss`.split.last.to_i; end
mm = 0

(1..51).each do |i|
  pid = Process.spawn(*ARGV.to_a, '1000')
  Thread.new do
    while true
      sleep 0.01
      m = mem(pid)
      mm = m if m > mm
    end
  end
  Process.waitall
  puts mm
  mm = 0
end

