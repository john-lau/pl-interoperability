#!/usr/bin/env ruby

pid = Process.spawn(*ARGV.to_a, '1000')

(1..50).each do |i|
  Process.spawn(*ARGV.to_a, '1000')
  Process.waitall
end

