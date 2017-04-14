#
# Just a file with songs.
#

class SongDatabase
  songs = []
  1.upto(9) do |i|
    s = {
      'file' => "#{i}.mp3",
      'Id' => i.to_s,
      'Last-Modified' => "2016-01-0#{i}",
      'Artist' => 'Fake-Artist',
      'AlbumArtist' => 'Fake-Album-Artist',
      'Title' => "Song number #{i}",
      'Album' => "Album number #{i}",
      'Track' => '1',
      'Date' => "200#{i}",
      'Genre' => 'Rock',
      'Time' => '240',
      'Pos' => i.to_s
    }.freeze
    songs << s
  end
  SONGS = songs.freeze

  def initialize
    @current_song_index = 0
  end

  def current_song
    SONGS[current_song_index]
  end

  def next_song
    @current_song_index += 1
    @current_song_index = 0 if @current_song_index == SONGS.size
    current_song
  end

  def previous_song
    @current_song_index -= 1
    @current_song_index = SONGS.size - 1 if @current_song_index.zero?
    current_song
  end

  def play_by_id(id)
    SONGS.each_with_index do |s, i|
      next unless s['Id'] == id
      @current_song_index = i
      return current_song
    end
    nil
  end

  def play_by_index(i)
    return unless SONGS[i]
    @current_song_index = i
    current_song
  end

  private

  attr_reader :current_song_index
end
