<?php
// comment here - is this right
class FileLister {
    public function listFiles() {
        return scandir(getcwd());
    }
}

/* Multiline 
as well
*/ 
$lister = new FileLister(); // yep
print_r($lister->listFiles()); /* nope */
?>

