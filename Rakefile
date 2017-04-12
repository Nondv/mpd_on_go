require_relative 'mpd/test/server/fake-mpd-server'

SERVER_PORT = 6789

task :default do
  sh 'go build'
end

task :test do
  Thread.abort_on_exception = true
  Thread.new { FakeMpdServer.new(SERVER_PORT).run }
  cd 'mpd'
  # it may cause troubles if server isn't running yet
  sh 'go test'
end
