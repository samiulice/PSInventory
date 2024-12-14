!include "LogicLib.nsh"
!include "FileFunc.nsh"
!insertmacro GetTempFileName

Section "Get Motherboard and CPU Info"
    ; Temporary file to store WMIC output
    ${GetTempFileName} $R0
    ; Get Motherboard Serial Number
    ExecDos::exec '"wmic baseboard get serialnumber"' "" "$R0"
    ClearErrors
    ReadINIStr $R1 "$R0" "Field" "1"  ; First line (skip header)
    ${If} ${Errors}
        MessageBox MB_OK "Failed to retrieve Motherboard ID."
    ${Else}
        MessageBox MB_OK "Motherboard Serial Number: $R1"
    ${EndIf}
    
    ; Get CPU ID
    ExecDos::exec '"wmic cpu get processorid"' "" "$R0"
    ClearErrors
    ReadINIStr $R2 "$R0" "Field" "1"  ; First line (skip header)
    ${If} ${Errors}
        MessageBox MB_OK "Failed to retrieve CPU ID."
    ${Else}
        MessageBox MB_OK "CPU ID: $R2"
    ${EndIf}
    
    ; Clean up
    Delete $R0
SectionEnd
