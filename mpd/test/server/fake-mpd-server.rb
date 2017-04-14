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
      @volume = 100
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

    attr_reader :song_db, :socket, :volume

    def execute_command(command)
      case command.chomp
      when 'status' then status
      when 'currentsong' then current_song
      when 'next' then next_song
      when 'previous' then previous_song
      when /setvol/ then setvol(command)
      when /play / then play(command)
      when /playid / then play_id(command)
      else raise "Undefinend command: #{command}"
      end
    end

    def setvol(full_command)
      argument = Integer(full_command.match(/setvol (.*)/)[1])
      raise(ArgumentError, 'volume should be 0..100') if argument < 0 || argument > 100
      @volume = argument
      "OK\n"
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
      song_db.next!
      "OK\n"
    end

    def previous_song
      song_db.previous!
      "OK\n"
    end

    def current_song
      hash_to_mpd_output(song_db.current) + "OK\n"
    end

    def status
      hash_to_mpd_output(status_hash) + "OK\n"
    end

    def status_hash
      { volume: volume,
        repeat: 0,
        random: 0,
        single: 0,
        consume: 0,
        playlist: 1,
        playlistlength: SongDatabase::SONGS.size,
        mixrampdb: '0.0',
        state: 'pause',
        song: song_db.current_song_index,
        songid: song_db.current['Id'],
        time: "0:#{song_db.current['Time']}",
        elapsed: '188.302',
        bitrate: '320',
        audio: '44100:24:2',
        nextsong: song_db.next_song_index,
        nextsongid: song_db.next_song['Id'] }
    end

    def hash_to_mpd_output(hash)
      hash.map { |k, v| "#{k}: #{v}\n" }.join
    end
  end
end
