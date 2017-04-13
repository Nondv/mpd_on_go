require 'socket'
require_relative 'songs'

class FakeMpdServer < TCPServer
  VERSION = 'fake-version'.freeze

  def run
    loop { Thread.start(ClientSession.new(accept), &:start) }
  end

  # This is neccessary to make server stateless between connections
  class ClientSession
    def initialize(socket)
      @socket = socket
      @current_song_index = 0
    end

    def start
      socket.puts "OK MPD #{VERSION}"
      until socket.closed?
        command = socket.gets
        next unless command
        response = execute_command(command)
        raise('response is not valid') unless response.is_a?(String) && response.end_with?("OK\n")
        socket.write(response)
      end
    rescue => e
      $stdout.puts("FAKE SERVER ERROR: #{e}")
      $stdout.puts(e.backtrace)
      socket.close
    end

    private

    attr_reader :current_song_index, :socket

    def execute_command(command)
      case command.chomp
      when 'currentsong' then current_song
      when 'next' then next_song
      when 'previous' then previous_song
      when /play / then play(command)
      when /playid / then play_id(command)
      else raise "Undefinend command: #{command}"
      end
    end

    def play_id(full_command)
      argument = full_command.match(/playid (.*)/)[1]

      # it is not obvious but if song is not found
      # session loop will raise error (as planned)
      SONGS.each_with_index do |s, i|
        next unless s['Id'] == argument
        @current_song_index = i
        return "OK\n"
      end
    end

    def play(full_comand)
      argument = full_comand.match(/play (.*)/)[1]
      @current_song_index = Integer(argument)
      "OK\n"
    rescue ArgumentError # from Integer()
      raise 'wrong argument'
    end

    def next_song
      inc_song_index
      "OK\n"
    end

    def previous_song
      dec_song_index
      "OK\n"
    end

    def current_song
      result = SONGS[current_song_index].map { |k, v| "#{k}: #{v}\n" }.join
      result << "OK\n"
      result
    end

    def inc_song_index
      @current_song_index += 1
      @current_song_index = 0 if @current_song_index == SONGS.size
    end

    def dec_song_index
      @current_song_index -= 1
      @current_song_index = SONGS.size - 1 if @current_song_index == SONGS.size
    end
  end
end
