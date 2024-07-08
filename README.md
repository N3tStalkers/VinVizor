# VIN Vizor

VIN Vizor is a Go program that retrieves and visualizes vehicle information using VIN (Vehicle Identification Number). It interacts with the NHTSA (National Highway Traffic Safety Administration) API to decode VINs and display detailed vehicle information in a formatted manner.

## Features

- **VIN Decoding**: Enter a VIN to retrieve detailed vehicle information.
- **Formatted Output**: Displays retrieved vehicle details in a visually appealing format.
- **Command-Line Interface**: Simple and intuitive CLI interface for ease of use.

## Installation

To run VIN Vizor, ensure you have Go installed on your system. Then, clone the repository and build the executable:

```bash
git clone https://github.com/N3tStalkers/VinVizor.git
cd VinVizor
go build
```

## Usage

Run the program and follow the prompts to enter a VIN:

```bash
./VinVizor
```

## Example

```plaintext
Enter VIN #: 1HGCM82633A123456

   #  #     #           #  #     #                        
   #  #                 #  #                              
   #  #    ##    ###    #  #    ##    ####    ##    # #   
   ####     #    #  #   ####     #      ##   #  #   ## #  
    ##      #    #  #    ##      #    ##     #  #   #     
    ##     ###   #  #    ##     ###   ####    ##    #     
                                                          
                                                          
VIN: 1HGCM82633A123456
Make: Honda
Model: Accord
Model Year: 2003
Body Class: Coupe
Drive Type: 
Engine Configuration: V=Shaped
Engine Cylinders: 6
Fuel Type: Gasoline
```

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.