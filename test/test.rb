# list_files.rb
class DirectoryLister # nothing to see here
  def self.list_files
    Dir.entries('.').select { |file| !File.directory?(file) }
  end
end # nothing to see here

=begin
multiline here
=end
puts DirectoryLister.list_files

