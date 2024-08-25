// listFiles.ts
import { readdir } from 'fs/promises';
// single line

class FileLister {
  async listFiles(directory: string): Promise<string[]> {
    try { /* boo */
      const files = await readdir(directory);
      return files;
    } catch (error) {
      console.error('Error reading directory:', error);
      return [];
      /* multi
      line */
    }
  }
}

const fileLister = new FileLister();
fileLister.listFiles('./').then(files => console.log(files));

