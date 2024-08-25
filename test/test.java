// ListFilesInDirectory.java
import java.io.File;

// Java comment here
public class ListFilesInDirectory {
    public static void main(String[] args) {
        File folder = new File(".");
        /* This
              is
              a
              multi-line
              comment */
        File[] listOfFiles = folder.listFiles();

        if (listOfFiles != null) {
            for (File file : listOfFiles) {
                if (file.isFile()) {
                    System.out.println("File: " + file.getName());
                } else if (file.isDirectory()) {
                    System.out.println("Directory: " + file.getName());
                }
            }
        } else {
            System.out.println("Unable to list files in the directory.");
        }
    }
} // done and done

