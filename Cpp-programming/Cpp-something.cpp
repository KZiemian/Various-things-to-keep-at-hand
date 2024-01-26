#include <iostream>
#include <fcntl.h>
#include <unistd.h>
#include <chrono>
#include <thread>


using namespace std;

int main() {
    const char* devicePath = "/dev/ttyUSB0";

    // Open the device file in read-only mode
    int fileDescriptor = open(devicePath, O_RDONLY);

    cout << "fileDescriptor = " << fileDescriptor << ".\n";

    if (fileDescriptor == -1) {
        // Unable to open the device file
        cerr << "Failed to open the device. Maby you forgot sudo? Try sudo ./name_of_file" << endl;
        return 1;
    }

    // Presenting the buffer to store the data.
    char buffer[256];

    fcntl(fileDescriptor, F_SETFL, O_NONBLOCK);

    // TRYING TO READ THE DATA FROM THE CPC not exactly knowing what is there
    ssize_t bytesRead = read(fileDescriptor, buffer, sizeof(buffer));

    if (bytesRead == -1) {
        // Error reading from the device
        cerr << "Error reading from the device." << endl;
    } else if (bytesRead == 0) {
        // No data read (device closed or reached EOF)
        cout << "No data available." << endl;
    } else {
        // Print the read data
        cout << "Read data: " << string(buffer, bytesRead) << endl;
    }

    // Close the file descriptor
    close(fileDescriptor);


    /* else {
        // Successfully opened the device file
        cout << "I am here." << endl;

        // Close the file descriptor
        close(fileDescriptor);
    }*/

    return 0;
}
