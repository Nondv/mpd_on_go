require 'socket'
require_relative 'songs'

class FakeMpdServer < TCPServer
  VERSION = 'fake-version'.freeze

  # TODO: use Thread.start(accept) ??
  def run
    loop do
      Thread.start(accept) do |socket|
        socket.puts "OK MPD #{VERSION}"
        command = socket.gets
        next unless command
        command.chop! # newline
        socket.write(execute_command(command))
      end
    end
  end

  private

  def execute_command(command)
    case command
    when 'currentsong' then currentsong
    else raise "Undefinend command: #{command}"
    end
  end

  def currentsong
    result = SONGS[0].map { |k, v| "#{k}: #{v}\n" }.join
    result << "OK\n"
    result
  end
end
