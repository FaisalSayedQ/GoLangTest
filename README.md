# GoLangTest

Learning and testing golang

## Setting up application in Local Machine

Pre requisites:
1. Install golang from https://golang.org/doc/install

To set up the application on your local box, follow the steps:

1. Clone or Download the github repository
2. `cd` to the downloaded app folder
3. Run the command `go run GoLangTest.go` or build and run through `go build && ./GoLangTest`
4. The application does not yet take settings from configuration file

## Testing the application

To test the application open http://localhost:8081/ in your browser

### Usage
1. Select currencies in both the drop downs for conversion
2. Type the amount in the top box
3. The conversion will happen when the input boxes/drop downs loose focus
4. Time displayed is local, calculated on each conversion however not formatted for optimal reading
5. If conversion from bottom currency is desired, remove the amount from top box

### Implementation Details
1. On the application start JSON data is fetched from fixer.io and parsed and stored in application memory
2. On accessing the application at http://localhost:8081/ The UI is loaded from a template
3. The UI sends data to the back end via http://localhost:8081/convert/IND/10/USD/0 url scheme
4. The URL is parsed and proper struct is loaded with the data, which is sent back as HTML to the AJAX call to be replaced in the template block

### Challenges
1. JSON parsing is not straight forward. Converting embedded objects to structs was particularly challenging
2. Template has potential however needs to be explored more
3. Parsing URL and converting numeric values to proper type needs to be explored

### To be done
1. Error handling
2. Edge cases handling
3. Test cases
4. Configurable settings