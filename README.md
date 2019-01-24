Just testing-learning rust, go. Some small benchmarks:

```
$ time php main.php 
It took 9 seconds
string(694) "{% block subject %}Czym jest Lorem Ipsum?{% endblock %}
{% block body %}Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.
Został po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach {{ years %}}. {{ century %}} w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker{% endblock %}"

real    0m9,113s
user    0m9,082s
sys     0m0,010s
```

----

```
$ time cargo run --release
    Finished release [optimized] target(s) in 0.02s
     Running `target/release/rust-test1`
It took time: Ok(1.007017123s)
[src/main.rs:40] result = "{% block subject %}Czym jest Lorem Ipsum?{% endblock %}\n{% block body %}Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.\nZostał po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach {{ years }}. {{ century }} w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker{% endblock %}\n"

real    0m1,050s
user    0m3,953s
sys     0m0,014s
```

-----

```
$ time ./main 
Goroutines amount 800 
Run took 4 seconds
{% subject block %}Czym jest Lorem Ipsum?{% endblock %}
{% body block %}Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.
Został po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach {{ years }}. {{ century }} w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker{% endblock %}

real    0m4,462s
user    0m15,557s
sys     0m1,598s
```

==================
CPU INFO:
```
$ lscpu
Architecture:        x86_64
CPU op-mode(s):      32-bit, 64-bit
Byte Order:          Little Endian
CPU(s):              4
On-line CPU(s) list: 0-3
Thread(s) per core:  2
Core(s) per socket:  2
Socket(s):           1
NUMA node(s):        1
Vendor ID:           GenuineIntel
CPU family:          6
Model:               42
Model name:          Intel(R) Core(TM) i5-2520M CPU @ 2.50GHz
Stepping:            7
CPU MHz:             840.991
CPU max MHz:         3200,0000
CPU min MHz:         800,0000
BogoMIPS:            4984.26
Virtualization:      VT-x
L1d cache:           32K
L1i cache:           32K
L2 cache:            256K
L3 cache:            3072K
NUMA node0 CPU(s):   0-3
Flags:               fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx rdtscp lm constant_tsc arch_perfmon pebs bts nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic popcnt tsc_deadline_timer aes xsave avx lahf_lm epb pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid xsaveopt dtherm ida arat pln pts flush_l1d

```

MEM INFO:
```
$ sudo dmidecode -t 17

# dmidecode 3.2
Getting SMBIOS data from sysfs.
SMBIOS 2.6 present.

Handle 0x0006, DMI type 17, 28 bytes
Memory Device
        Array Handle: 0x0005
        Error Information Handle: Not Provided
        Total Width: 64 bits
        Data Width: 64 bits
        Size: 4096 MB
        Form Factor: SODIMM
        Set: None
        Locator: ChannelA-DIMM0
        Bank Locator: BANK 0
        Type: DDR3
        Type Detail: Synchronous
        Speed: 1333 MT/s
        Manufacturer: Samsung
        Serial Number: 0111086D
        Asset Tag: 9876543210
        Part Number: M471B5273CH0-CH9
        Rank: Unknown

Handle 0x0007, DMI type 17, 28 bytes
Memory Device
        Array Handle: 0x0005
        Error Information Handle: Not Provided
        Total Width: 64 bits
        Data Width: 64 bits
        Size: 4096 MB
        Form Factor: SODIMM
        Set: None
        Locator: ChannelB-DIMM0
        Bank Locator: BANK 2
        Type: DDR3
        Type Detail: Synchronous
        Speed: 1333 MT/s
        Manufacturer: 075D
        Serial Number: 00000000
        Asset Tag: 9876543210
        Part Number: GR1333S364L9S/4G
        Rank: Unknown
```


