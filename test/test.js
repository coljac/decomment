const fs = require('fs');
const path = require('path');

class DirectoryLister {
  constructor(directoryPath = '.') {
    this.directoryPath = directoryPath;
  } // Commenty comment

  listFiles() {
    fs.readdir(this.directoryPath, (err, files) => {
      if (err) {
        console.error('Error reading directory:', err);
        return;
      }
      files.forEach(file => console.log(file));
    });
  }
} // Commenty comment

/* Why are you reading this
  * You should be reading the code above
  * Not this comment
  */
const lister = new DirectoryLister();
lister.listFiles();

