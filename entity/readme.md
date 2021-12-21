# Entity 

An entity is any singular, identifiable and separate object.

Here we have 4 main types of entities:
 - Parking.go contains information about the parking spots and groups of them. Parking price is applied at parking group level to all parking spots in the group, but each parking spot _might_ have individual pricing model
 - Municipality.go - parking groups belongs to district -> city -> state -> country. 
 - Admin.go - someone should manage the parking spots and groups. Administrator and Inspectors managed in this file 
 - User.go - contains information about end users of the application. User can book parking spot, check available and get timer for the parking spot 


## Reference:
[Property Name Format](https://google.github.io/styleguide/jsoncstyleguide.xml?showone=Property_Name_Format#Property_Name_Format): 
