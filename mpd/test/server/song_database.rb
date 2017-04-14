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

  attr_reader :current_song_index

  def initialize
    @current_song_index = 0
  end

  def current
    SONGS[current_song_index]
  end

  def next!
    @current_song_index = next_song_index
    current
  end

  def previous!
    @current_song_index = previous_song_index
    current
  end

  def next_song
    SONGS[next_song_index]
  end

  def previous_song
    SONGS[previous_song]
  end

  def next_song_index
    res = @current_song_index + 1
    res == SONGS.size ? 0 : res
  end

  def previous_song_index
    res = @current_song_index - 1
    res.zero? ? SONGS.size - 1 : res
  end

  def play_by_id(id)
    SONGS.each_with_index do |s, i|
      next unless s['Id'] == id
      @current_song_index = i
      return current
    end
    nil
  end

  def play_by_index(i)
    return unless SONGS[i]
    @current_song_index = i
    current
  end
end
