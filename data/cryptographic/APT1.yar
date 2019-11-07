rule APT_Bouncer : APT
 {
 meta:
     internal = true
     test = false
     score = 100
     threat = "PMA (APT Bouncer)"
     behaviour_class = "Family"
     author = "Adam"
     date = "2013-08-21"
     description = "APT Bouncer"
     version = 1
 strings:
     $string0 = "do_pivot:"
     $string1 = "sqlpass.dic"
     $string2 = "sa:1qaz2wsx"
 condition:
     all of them
 }

 rule ComputraceAgent
 {
 meta:
     internal = true
     test = false
     score = 40
     threat = "PUA (Computrace Agent)"
     behaviour_class = "Family"
     author = "Kaspersy Lab"
     date = "2014-09-23"
     description = "Absolute Computrace Agent Executable"
     version     = 1
 strings:
     $a = {D1 E0 F5 8B 4D 0C 83 D1 00 8B EC FF 33 83 C3 04}
     $mz = {4d 5a}
     $b1 = {72 70 63 6E 65 74 70 2E 65 78 65 00 72 70 63 6E 65 74 70 00}
     $b2 = {54 61 67 49 64 00}
 condition:
     ($mz at 0 ) and ($a or ($b1 and $b2))
 }

 rule ONHAT : ONHAT
 {
     meta:
         internal     = true
         test         = false
         score        = 75
         threat       = "PMA (ONHAT Proxy Tool)"
         threat_class = "security-tool"
         behaviour_class = "Family"
         author       = "Alex"
         date         = "2014-10-13"
         description  = "ONHAT"
         version      = 1

         strings:
                 $ = "[ONHAT] ACCEPTS ERROR."
                 $ = "[ONHAT] CREATES EVENT (%d, %d)."
                 $ = "[ONHAT] LISTENS (%d.%d.%d.%d, %d) ERROR."
         condition:
                 all of them
 }
 rule Aumlib : Aumlib
 {
     meta:
         internal     = true
         test         = true
         score        = 100
         threat       = "PMA (Backdoor Aumlib)"
         threat_class = "trojan"
         behaviour_class = "Family"
         author       = "Alex"
         date         = "2014-11-17"
         description  = "Aumlib"
         version      = 1
     strings:
         $a = "/bbs/info.asp"
         $b = "aumLib.ini"

     condition:
         all of them
 }

 rule Duqu: Duqu
 {
 meta:
     internal        = true
     test            = true
     score           = 75
     threat          = "PMA ( Duqu Driver )"
     threat_class    = "rootkit"
     author          = "Alexander Hanel"
     date            = "11/24/2014"
     description     = "PMA ( Duqu Driver )"
     version         = 1

 strings:
     $ = "\\Device\\Gpd1" wide
     $ = "\\Device\\Gpd1" wide
     $ = "KdDebuggerEnabled"
     $ = "InitSafeBootMode"
     $ = {8B 0D D4 33 01 00 83 39  00 74 07 B8 01 00 00 C0}

 condition:
     all of them
 }

 rule regin: regin
 {
     meta:
         internal        = true
         test            = true
         score           = 75
         threat          = "PMA ( Regin Driver )"
         threat_class    = "rootkit"
         author          = "Alexander Hanel"
         date            = "11/24/2014"
         description     = "PMA ( Regin Driver )"
         version         = 1

     strings:
         // The bytes are the encoded registry keys
         $a = {6F 43 52 AE 6B 4B 6A F2 62 45 52 80 49 78 5F C6 46 5D 54 80 41 7C 4B DD 77 4A 7A}
         $b = "NtBuildNumber"
     condition:
         all of them
 }
