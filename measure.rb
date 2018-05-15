#!/usr/bin/env ruby

def mem(pid); `ps p #{pid} -o rss`.split.last.to_i; end
def cpu(pid); `ps p #{pid} -o cpu`.split.last.to_i; end
t = Time.now
pid = Process.spawn(*ARGV.to_a, '1000')
mm = 0
mcpu = 0

Thread.new do
  mm = mem(pid)
  mcpu = cpu(pid)
  while true
    sleep 0.1
    m = mem(pid)
    cpu = mem(pid)
    mm = m if m > mm
    mcpu = cpu if cpu > mcpu
  end
end

Process.waitall
STDERR.puts "%.2fs, %.1fMb, %f CPU" % [Time.now - t, mm / 1024.0, mcpu]

