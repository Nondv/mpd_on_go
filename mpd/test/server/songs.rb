#
# Just a file with songs.
#

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
