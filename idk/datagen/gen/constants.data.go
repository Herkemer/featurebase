package gen

var Animals = []string{"Millipede", "Reindeer", "Otter", "Umbrellabird", "Bandicoot", "Macaroni Penguin", "Australian Cattle Dog", "Hercules Beetle", "Magpie", "Bernese Mountain Dog",
	"Royal Penguin", "Manatee", "Black Russian Terrier", "Badger", "Indian Elephant", "Indochinese Tiger", "Centipede", "Bluetick Coonhound", "Humpback Whale", "Caiman Lizard",
	"Grasshopper", "Tortoise", "Alpine Dachsbracke", "Quetzal", "Aldabra Giant Tortoise", "Electric Eel", "Vervet Monkey", "Mountain Lion", "Geoffroys Tamarin", "Collared Peccary",
	"Doberman Pinscher", "Megalodon", "Insect", "Shiba Inu", "Tiger Shark", "Aye Aye", "Black Rhinoceros", "Gar", "Stick Insect", "Marine Toad",
	"Grouse", "Javanese", "Water Dragon", "Whippet", "Giant Clam", "Caiman", "Fennec Fox", "Weasel", "Rat", "Golden Retriever Complete Pet Guide",
	"Starfish", "Moth", "Horned Frog", "Water Vole", "King Crab", "Common Frog", "Neanderthal", "Fossa", "Frigatebird", "Spider Monkey",
	"Appenzeller Dog", "Emperor Penguin", "Saint Bernard", "Eagle", "Platypus", "Monkfish", "Lizard", "Barracuda", "Wombat", "Pika",
	"Yak", "Ibis", "Beetle", "Horse", "Arctic Fox", "Chihuahua", "African Wild Dog", "Bat", "Dogue De Bordeaux", "Akita",
	"Mastiff", "Fox Terrier", "Thorny Devil", "Dachshund", "Zonkey", "Dormouse", "Guinea Pig", "Edible Frog", "Barnacle", "Asian Palm Civet",
	"Irish WolfHound", "Giant Panda Bear", "Abyssinian", "Toucan", "Wild Boar", "Grizzly Bear", "Borneo Elephant", "Dusky Dolphin", "Gecko", "Golden Lion Tamarin",
	"Deer", "Indri", "Stoat", "Skunk", "Sea Lion", "Turkey", "Octopus", "Pelican", "Meerkat", "Pig",
	"Glass Lizard", "Fur Seal", "Pekingese", "Bottlenose Dolphin", "Fox", "Siamese", "Scorpion Fish", "Howler Monkey", "Anatolian Shepherd Dog", "Mountain Gorilla",
	"Burrowing Frog", "African Bush Elephant", "Havanese", "Sumatran Elephant", "Glow Worm", "Saber-Toothed Tiger", "Japanese Chin", "Tuatara", "Estrela Mountain Dog", "Cow",
	"Dwarf Crocodile", "Flounder", "Tetra", "Gorilla", "Bearded Collie", "Angelfish", "Gentoo Penguin", "Afghan Hound", "Porpoise", "Australian Kelpie Dog",
	"Pygmy Marmoset", "Harpy Eagle", "Ostrich", "Woolly Mammoth", "Duck", "Binturong", "Sea Slug", "Sumatran Tiger", "Cassowary", "King Penguin",
	"Bongo", "Basset Hound", "Siberian Husky", "Camel Spider", "Boxer Dog", "Cockroach", "Antelope", "Foxhound", "Beagle", "Impala",
	"Snapping Turtle", "Grey Reef Shark", "Jellyfish", "Pheasant", "Cavalier King Charles Spaniel", "Monarch Butterfly", "Wrasse", "Drever", "Catfish", "Malayan Civet",
	"Bearded Dragon", "Red Wolf", "Cuttlefish", "Chipmunk", "English Springer Spaniel", "Tiger Salamander", "Epagneul Pont Audemer", "Mongoose", "Somali", "Parrot",
	"Burmese", "Greyhound", "Tiger", "Tree Frog", "Skate Fish", "Flamingo", "Bonobo", "Ibizan Hound", "Spadefoot Toad", "Japanese Macaque",
	"Leopard Seal", "Leopard Tortoise", "Moose", "Chicken", "Avocet", "Australian Mist", "Little Penguin", "Quail", "Elephant Seal", "English Cocker Spaniel",
	"Humboldt Penguin", "Bedlington Terrier", "Blue Whale", "Tibetan Mastiff", "Pink Fairy Armadillo", "Hedgehog", "Bavarian Mountain Hound", "Basking Shark", "Quoll", "West Highland Terrier",
	"Polar Bear", "Egyptian Mau", "Fish", "Crocodile", "Vulture", "Arctic Wolf", "Javan Rhinoceros", "Capybara", "Baboon", "Gharial",
	"Hawaiian Crow", "Piranha", "Australian Shepherd", "Fishing Cat", "Sponge", "Guppy", "Tasmanian Devil", "Chesapeake Bay Retriever", "Lobster", "Elephant",
	"Labradoodle – The Complete Guide For Owners", "Radiated Tortoise", "Dugong", "French Bulldog", "Slow Worm", "Bulldog", "Giant African Land Snail", "African Tree Toad", "Serval", "Nurse Shark",
	"Kangaroo", "Jackal", "Neapolitan Mastiff", "Cockatoo", "Pond Skater", "Okapi", "Hyena", "Border Terrier", "Silver Dollar", "Dog",
	"Coati", "Butterfly", "Moorhen", "Numbat", "Human", "King Cobra", "Green Bee-Eater", "Common Loon", "Cross River Gorilla", "Moray Eel",
	"Bloodhound", "Spixs Macaw", "Sperm Whale", "Russian Blue", "Persian", "Chimpanzee", "Tiffany", "African Civet", "Koala", "Tapir",
	"Minke Whale", "Clumber Spaniel", "Leopard", "Liger", "Red Panda", "Wolf Spider", "Mayfly", "Zebra", "Bombay", "Rattlesnake",
	"African Forest Elephant", "Yorkshire Terrier", "Jack Russel", "Maltese", "Booby", "Bison", "Woodpecker", "Squid", "Uguisu", "Sumatran Orang-utan",
	"Asian Giant Hornet", "Pink Dolphin", "German Pinscher", "Dalmatian", "Sun Bear", "Caracal", "Camel", "Vampire Bat", "Mole", "Dolphin",
	"Nightingale", "German Shepherd Guide", "Coyote", "Mandrill", "Newfoundland", "Cairn Terrier", "Kingfisher", "Darwin’s Frog", "Mouse", "Aardvark",
	"Greater Swiss Mountain Dog", "Pied Tamarin", "Monkey", "Raccoon", "Sea Turtle", "Sloth", "Airedale Terrier", "Alligator", "Orang-utan", "Wyoming Toad",
	"Shrimp", "Steller’s Sea Cow", "Cotton-top Tamarin", "Asiatic Black Bear", "Western Lowland Gorilla", "Alaskan Malamute", "Caterpillar", "Spiny Dogfish", "Bull Terrier", "Maned Wolf",
	"Cocker Spaniel", "Buffalo", "Lemming", "Crane", "Tang", "Monitor Lizard", "Crab-Eating Macaque", "Kudu", "Dunker", "Border Collie",
	"Bobcat", "Grey Mouse Lemur", "South China Tiger", "Bolognese Dog", "Kiwi", "Horseshoe Crab", "Highland Cattle", "Stingray", "Macaw", "Patas Monkey",
	"Bornean Orang-utan", "Birman", "Emu", "Spectacled Bear", "Lionfish", "Golden Masked Owl", "Honey Badger", "Wildebeest", "Quokka", "Hippopotamus",
	"Heron", "Cichlid", "Water Buffalo", "Chinook", "African Penguin", "Common Toad", "Termite", "Affenpinscher", "Leaf-Tailed Gecko", "Maine Coon",
	"Lynx", "Bactrian Camel", "Water Spaniel", "Harrier", "Manta Ray", "Coral", "Dogo Argentino", "Gerbil", "Australian Terrier", "Masked Palm Civet",
	"Snake", "Desert Tortoise", "Axolotl", "Greenland Dog", "Indian Palm Squirrel", "Fluke Fish (summer flounder)", "Woolly Monkey", "Ladybug", "Wolf", "Common Buzzard",
	"African Palm Civet", "Cesky Fousek", "Mule", "Amur Leopard", "Norwegian Forest", "Flying Squirrel", "Eskimo Dog", "Monte Iberia Eleuth", "Lion", "Eastern Lowland Gorilla",
	"Blobfish", "Chamois", "Arctic Hare", "Grey Seal", "English Shepherd", "Swan", "Sea Urchin", "Birds Of Paradise", "Butterfly Fish", "Blue Jay",
	"Staffordshire Bull Terrier", "Siamese Fighting Fish", "White-Faced Capuchin", "Whale Shark", "Snail", "Pit Bull Terrier", "Squirrel Monkey", "Hoopoe", "Sea Dragon", "Seahorse",
	"Great Dane", "Clown Fish", "Snowy Owl", "Anteater", "White Rhinoceros", "Boykin Spaniel", "Green Anole", "Pug", "Dhole", "Goose",
	"Deutsche Bracke", "Banded Palm Civet", "Llama", "Sri Lankan Elephant", "Welsh Corgi: The Complete Pet Guide", "Tropicbird", "Rhinoceros", "Curly Coated Retriever", "Leopard Cat", "Black Widow Spider",
	"Ragdoll", "Horn Shark", "Donkey", "Elephant Shrew", "Armadillo", "Keel Billed Toucan", "Purple Emperor", "Scorpion", "Chameleon", "Sea Otter",
	"Giant Schnauzer", "Bear", "Honey Bee", "Zebu", "African Clawed Frog", "Blue Lacy Dog", "Pangolin", "Cat", "Collie", "Dodo",
	"Brown Bear", "Akbash", "Tapanuli Orang-utan", "Puma", "Iguana", "Uakari", "Chinese Crested Dog", "Magellanic Penguin", "Hare", "Molly",
	"Peacock", "Saola", "Squirrel", "Basenji Dog", "Bumblebee", "Ferret", "Staffordshire Terrier", "Ainu Dog", "Rottweiler", "Beaver",
	"Sparrow", "Sea Squirt", "Labrador Retriever", "Pool Frog", "Bichon Frise", "Oyster", "Panther", "Hammerhead Shark", "Frog", "Ocelot",
	"Coonhound", "Rabbit", "Adelie Penguin", "Olm", "Porcupine", "Hummingbird", "Wasp", "Tawny Owl", "Walrus", "X-Ray Tetra",
	"Goat", "Komodo Dragon", "Fly", "Western Gorilla", "Golden-Crowned Flying Fox", "Stag Beetle", "Budgerigar", "Chinchilla", "Finnish Spitz", "Woodlouse",
	"Bull Mastiff", "Jaguar", "River Turtle", "Wolverine", "Siberian Tiger", "Jerboa", "Crab", "Emperor Tamarin", "Kakapo", "Hermit Crab",
	"Eskimo Dog", "North American Black Bear", "Galapagos Penguin", "Chow Chow", "Red-handed Tamarin", "Albatross", "Roseate Spoonbill", "Siberian", "Striped Rocket Frog", "Malayan Tiger",
	"Boston Terrier: Complete Pet Guide", "Poodle", "Clouded Leopard", "Red Knee Tarantula", "Yellow-Eyed Penguin", "Zebra Shark", "Lemur", "Raccoon Dog", "Pointer", "Newt",
	"Dragonfly", "Carolina Dog", "Xerus", "Crested Penguin", "Warthog", "Ermine", "Frilled Lizard", "Field Spaniel", "Himalayan", "Turkish Angora",
	"Discus", "Echidna", "Markhor", "Pygmy Hippopotamus", "Fin Whale", "Canaan Dog", "Marsh Frog", "Gopher", "Bull Shark", "Salamander",
	"Brazilian Terrier", "Balinese", "Asian Elephant", "Prawn", "Robin", "Entlebucher Mountain Dog", "Snowshoe", "Galapagos Tortoise", "Barn Owl", "Proboscis Monkey",
	"Wallaby", "Gila Monster", "Indian Rhinoceros", "Bullfrog", "Norfolk Terrier", "Sumatran Rhinoceros", "Bengal Tiger", "Penguin", "Mongrel", "Cougar",
	"Narwhal", "Old English Sheepdog", "Flat Coat Retriever", "Pere Davids Deer", "Eastern Gorilla", "Barb", "Golden Oriole", "Pademelon", "Sheep", "Indian Star Tortoise",
	"Aurochs", "Possum", "Zorse", "Chinstrap Penguin", "Sucker Fish", "Great White Shark", "Gibbon", "Guinea Fowl", "Sand Lizard", "Earwig",
	"Seal", "Irish Setter Complete Pet Guide", "Rock Hyrax", "Cheetah", "Puffin", "American Bulldog", "Pufferfish", "White Tiger", "Hamster", "Rockhopper Penguin",
	"Shih Tzu", "Fire-Bellied Toad", "Cuscus", "Pike", "Poison Dart Frog", "Dingo", "Puss Moth", "Ant", "Tarsier", "Scimitar-horned Oryx",
	"Long-Eared Owl", "Falcon", "Killer Whale", "Bird", "Giraffe", "Opossum"}
