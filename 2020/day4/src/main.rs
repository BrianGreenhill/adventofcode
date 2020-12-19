use std::env;
use std::fs;
use regex::Regex;

#[derive(Debug, PartialEq)]
struct Passport {
    byr: Option<usize>,
    iyr: Option<usize>,
    eyr: Option<usize>,
    hgt: Option<Height>,
    hgt_str: Option<String>,
    hcl: Option<String>,
    ecl: Option<String>,
    pid: Option<String>,
    cid: Option<String>,
}

impl Passport {
    fn new() -> Passport {
        Passport {
            byr: None,
            iyr: None,
            eyr: None,
            hgt: None,
            hgt_str: None,
            hcl: None,
            ecl: None,
            pid: None,
            cid: None,
        }
    }
    fn has_fields(&self) -> bool {
        self.byr.is_some()
            && self.iyr.is_some()
            && self.eyr.is_some()
            && self.hgt_str.is_some()
            && self.hcl.is_some()
            && self.ecl.is_some()
            && self.pid.is_some()
    }

    fn is_valid(&self) -> bool {
        self.valid_byr()
            && self.valid_iyr()
            && self.valid_eyr()
            && self.valid_hgt()
            && self.valid_hcl()
            && self.valid_ecl()
            && self.valid_pid()
    }

    fn valid_pid(&self) -> bool {
        let re = Regex::new("^[0-9]{9}$").expect("Failed to build regex");
        match &self.pid {
            None => false,
            Some(p) => re.is_match(&p.as_str())
        }
    }

    fn valid_ecl(&self) -> bool {
        let valid_colors = vec!["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];
        match &self.ecl {
            None => false,
            Some(c) => valid_colors.contains(&c.as_str()),
        }
    }

    fn valid_hcl(&self) -> bool {
        let re = Regex::new("^#[0-9a-f]{6}$")
            .expect("Failed to make regex");
        match &self.hcl {
            None => false,
            Some(hair) => re.is_match(hair.as_str()),
        }
    }
    fn valid_byr(&self) -> bool {
        match self.byr {
            None => false,
            Some(n) => n >= 1920 && n <= 2002,
        }
    }

    fn valid_iyr(&self) -> bool {
        match self.iyr {
            None => false,
            Some(n) => n >= 2010 && n <= 2020,
        }
    }

    fn valid_eyr(&self) -> bool {
        match self.eyr {
            None => false,
            Some(n) => n >= 2020 && n <= 2030,
        }
    }

    fn valid_hgt(&self) -> bool {
        match &self.hgt {
            None => false,
            Some(h) => h.is_valid(),
        }
    }
}

#[derive(Debug, PartialEq)]
struct Height {
    measure: usize,
    unit: String,
}

impl Height {
    fn parse(hgt_str: &str) -> Option<Height> {
        let re = Regex::new("(\\d+)(in|cm)").expect("Unable to create Regex");
        match re.captures(hgt_str) {
            None => None,
            Some(captures) => {
                let h = Height {
                    measure: str::parse(captures.get(1).unwrap().as_str())
                        .expect("Unable to parse number"),
                    unit: String::from(captures.get(2).unwrap().as_str()),
                };
                Some(h)
            }
        }
    }
    fn is_valid(&self) -> bool {
        match self.unit.as_str() {
            "cm" => self.measure >= 150 && self.measure <= 193,
            "in" => self.measure >= 59 && self.measure <= 76,
            _ => panic!("Not a valid unit!"),
        }
    }
}

fn parse_passport(passport_str: &str) -> Passport {
    let kv: Vec<&str> = passport_str
        .lines()
        .flat_map(|line| line.split(" "))
        .collect();
    let mut pass = Passport::new();
    for key_val in kv {
        let pair: Vec<&str> = key_val.split(":").collect();
        match *(pair.get(0).unwrap()) {
         "byr" =>  pass.byr = Some(str::parse(*pair.get(1).unwrap()).unwrap()),
         "iyr" => pass.iyr = Some(str::parse(*pair.get(1).unwrap()).unwrap()),
         "eyr" => pass.eyr = Some(str::parse(*pair.get(1).unwrap()).unwrap()),
         "hgt" =>  {
            pass.hgt_str = Some(str::parse(*pair.get(1).unwrap()).unwrap());
            pass.hgt = Height::parse(*pair.get(1).unwrap());
         }
         "hcl" => pass.hcl = Some(String::from(*pair.get(1).unwrap())),
         "ecl" => pass.ecl = Some(String::from(*pair.get(1).unwrap())),
         "pid" => pass.pid = Some(String::from(*pair.get(1).unwrap())),
         "cid" => pass.cid = Some(String::from(*pair.get(1).unwrap())),
         _ => panic!("Found passport code that doesn't match"),
        }
    }
    pass
}

fn parse_input(input: &str) -> Vec<Passport> {
    input
        .split("\n\n")
        .map(|passport_str| parse_passport(passport_str))
        .collect()
}

fn count_valid_passports(input: &Vec<Passport>) -> usize {
    input
        .iter()
        .filter(|pass| pass.has_fields())
        .count()
}

fn count_valid_passports_part_two(input: &Vec<Passport>) -> usize {
    input
        .iter()
        .filter(|pass| pass.is_valid())
        .count()
}

fn main() {
    // accept data
    let args: Vec<String> = env::args().collect();
    let filename = &args[1];
    let contents = get_contents(filename);
    let passports = parse_input(&contents);
    // validate passport fields
    println!("Valid passports: {}", count_valid_passports(&passports));
    println!("Valid passports (Part 2): {}", count_valid_passports_part_two(&passports));
}

fn get_contents(filename: &String) -> String {
    let contents = fs::read_to_string(filename)
        .expect("Could not read file");
    contents
}
