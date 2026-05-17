# wordlist
Combined wordlist of 210000+ unique entries build from different sources.

## Sources
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/100k-most-used-passwords-NCSC.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2020-200_most_used_passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2023-200_most_used_passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2024-197_most_used_passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2025-199_most_used_passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/500-worst-passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/French-common-password-list-top-20000.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/German_common-password-list-top-100000.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/Spanish_1000-common-usernames-and-passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/common-passwords-win.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/medical-devices.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/top-20-common-SSH-passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/top-passwords-shortlist.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Cracked-Hashes/milw0rm-dictionary.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/Malware/conficker.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/corporate_passwords.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/seasons.txt
* https://github.com/danielmiessler/SecLists/blob/master/Passwords/stupid-ones-in-production.txt

## Building
```
xz -k -z -e -9 wordlist.txt
```

## License
Released under the [MIT License](LICENSE.md).