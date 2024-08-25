#!/usr/bin/perl
use strict;
use warnings;
# This is a comment

opendir(my $dh, ".") or die "Can't open directory: $!";
while (my $file = readdir($dh)) {
    next if $file eq '.' or $file eq '..';
    print "$file\n";
}
=for comment

Example of multiline comment.
# This should stay
Example of multiline comment.

=cut

closedir($dh);
