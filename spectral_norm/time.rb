#!/usr/bin/env ruby

# def mem(pid); `ps p #{pid} -o rss`.split.last.to_i; end
# def cpu(pid); `ps p #{pid} -o pcpu`.split.last.to_i; end

# t = Time.now
pid = Process.spawn(*ARGV.to_a, '1000')
# mem_arr = []
# cpu_arr = []

(1..50).each do |i|
  Process.spawn(*ARGV.to_a, '1000')
  Process.waitall
end

# Thread.new do
#   mem_arr << mem(pid)
#   cpu_arr << cpu(pid)
#   while true
#     sleep 0.1
#     m = mem(pid)
#     mem_arr << m
#     cpu = cpu(pid)
#     cpu_arr << cpu
#     # mm = m if m > mm
#     # mcpu = cpu if cpu > mcpu
#   end
# end

# Process.waitall

# min_mem = mem_arr.min
# max_mem = mem_arr.max
# total_mem = mem_arr.inject(:+)
# len_mem = mem_arr.length
# avg_mem = total_mem.to_f/len_mem

# min_cpu = cpu_arr.min
# max_cpu = cpu_arr.max
# total_cpu = cpu_arr.inject(:+)
# len_cpu = cpu_arr.length
# avg_cpu = total_cpu.to_f/len_cpu

# STDERR.puts "TIME: %.2fs" % [Time.now - t]
# STDERR.puts "MEM: min %fkB, max %fkB, avg %fkB" % [min_mem, max_mem, avg_mem]
# STDERR.puts "CPU: min %f, max %f, avg %f" % [min_cpu, max_cpu, avg_cpu]

