# Dawson Butts
# Feb 21st 

README for main folder

In the main folder, I altered main() to accept another command line input
so that facets could be input and used in host.go. I changed the input 
validation to include 3 inputs (main, query, facets). 

I also added the new parameter to the call of the HostSearch method 
further down into the main function. 

This method can be called on the command line by: 
SHODAN_API_KEY=Your_Api_Key ./main <searchTerm> <facets>
