# wordlist
Combined wordlist of 210000+ unique entries build from different sources.

## Sources
* [English Common Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/100k-most-used-passwords-NCSC.txt)
* [French Common Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/French-common-password-list-top-20000.txt)
* [German Common Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/German_common-password-list-top-100000.txt)
* [Spanish Common Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/Language-Specific/Spanish_1000-common-usernames-and-passwords.txt)
* [2025 Most Used Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2025-199_most_used_passwords.txt)
* [2024 Most Used Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2024-197_most_used_passwords.txt)
* [2023 Most Used Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2023-200_most_used_passwords.txt)
* [2020 Most Used Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/2020-200_most_used_passwords.txt)
* [Most Common Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt)
* [Common Passwords Shortlist](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/top-passwords-shortlist.txt)
* [Common Win Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/common-passwords-win.txt)
* [Common SSH Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/top-20-common-SSH-passwords.txt)
* [Worst Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/500-worst-passwords.txt)
* [Milw0rm Dictionary](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Cracked-Hashes/milw0rm-dictionary.txt)
* [Conficker Dictionary](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Malware/conficker.txt)
* [Corporate Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/corporate_passwords.txt)
* [Production Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/stupid-ones-in-production.txt)
* [Medical Device Passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/medical-devices.txt)
* [Seasons](https://github.com/danielmiessler/SecLists/blob/master/Passwords/seasons.txt)

## Building
```
xz -k -z -e -9 wordlist.txt
```

## License
Released under the [MIT License](LICENSE.md).