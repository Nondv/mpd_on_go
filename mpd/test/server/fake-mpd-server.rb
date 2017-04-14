require 'socket'
require_relative 'song_database'

class FakeMpdServer < TCPServer
  VERSION = 'fake-version'.freeze

  def run
    loop { Thread.start(ClientSession.new(accept), &:start) }
  end

  # This is neccessary to make server stateless between connections
  class ClientSession
    def initialize(socket)
      @socket = socket
      @song_db = SongDatabase.new
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

    attr_reader :song_db, :socket

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
      song_db.play_by_id(argument) && "OK\n"
    end

    def play(full_comand)
      argument = full_comand.match(/play (.*)/)[1]
      song_db.play_by_index(Integer(argument)) && "OK\n"
    rescue ArgumentError # from Integer()
      raise 'wrong argument'
    end

    def next_song
      song_db.next_song
      "OK\n"
    end

    def previous_song
      song_db.previous_song
      "OK\n"
    end

    def current_song
      result = song_db.current_song.map { |k, v| "#{k}: #{v}\n" }.join
      result << "OK\n"
      result
    end
  end
end
