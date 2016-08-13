package gogame

import "github.com/veandco/go-sdl2/sdl"

// Input combines all of the input methods together.
type Input interface {
	WindowInput
	MouseInput
	KeyboardInput
}

// WindowInput gets input from a window.
type WindowInput interface {
	// WindowPosition returns the position of the window on the screen in pixels.
	WindowPosition() (x, y int)

	// WindowSize returns the size of the window in pixels.
	WindowSize() (w, h int)

	// WindowMoved checks if the window has just been moved.
	WindowMoved() bool

	// WindowResized checks if the window has just been resized.
	WindowResized() bool

	// WindowClosed checks if the window has just been closed (X, Alt+F4, etc.).
	WindowClosed() bool

	// WindowHasFocus checks if the window has focus.
	WindowHasFocus() bool

	// WindowLostFocus checks if the window has just lost focus.
	WindowLostFocus() bool

	// WindowGainedFocus checks if the window has just gained focus.
	WindowGainedFocus() bool
}

// MouseInput gets input from a mouse device.
type MouseInput interface {
	// MousePosition returns the position of the mouse relative to the window.
	MousePosition() (x, y int)

	// MouseDelta returns the difference between mouse's current and previous position.
	MouseDelta() (dx, dy int)

	// MouseDown checks if a mouse button is currently pressed down.
	MouseDown(button int) bool

	// MouseJustDown checks if a mouse button has just been pressed down.
	MouseJustDown(button int) bool
}

// KeyboardInput gets input from a keyboard device.
type KeyboardInput interface {
	// KeyDown checks if a key is currently pressed down.
	KeyDown(key int) bool

	// KeyJustDown checks if a key has just been pressed down.
	KeyJustDown(key int) bool
}

// Enumeration of all mouse buttons.
const (
	MouseButtonLeft   = sdl.BUTTON_LEFT
	MouseButtonMiddle = sdl.BUTTON_MIDDLE
	MouseButtonRight  = sdl.BUTTON_RIGHT
	MouseButtonX1     = sdl.BUTTON_X1
	MouseButtonX2     = sdl.BUTTON_X2
)

// Enumeration of all keyboard keys.
const (
	KeyUnknown = sdl.K_UNKNOWN

	KeyReturn     = sdl.K_RETURN
	KeyEscape     = sdl.K_ESCAPE
	KeyBackspace  = sdl.K_BACKSPACE
	KeyTab        = sdl.K_TAB
	KeySpace      = sdl.K_SPACE
	KeyExclaim    = sdl.K_EXCLAIM
	KeyQuoteDbl   = sdl.K_QUOTEDBL
	KeyHash       = sdl.K_HASH
	KeyPercent    = sdl.K_PERCENT
	KeyDollar     = sdl.K_DOLLAR
	KeyAmpersand  = sdl.K_AMPERSAND
	KeyQuote      = sdl.K_QUOTE
	KeyLeftParen  = sdl.K_LEFTPAREN
	KeyRightParen = sdl.K_RIGHTPAREN
	KeyAsterisk   = sdl.K_ASTERISK
	KeyPlus       = sdl.K_PLUS
	KeyComma      = sdl.K_COMMA
	KeyMinus      = sdl.K_MINUS
	KeyPeriod     = sdl.K_PERIOD
	KeySlash      = sdl.K_SLASH
	Key0          = sdl.K_0
	Key1          = sdl.K_1
	Key2          = sdl.K_2
	Key3          = sdl.K_3
	Key4          = sdl.K_4
	Key5          = sdl.K_5
	Key6          = sdl.K_6
	Key7          = sdl.K_7
	Key8          = sdl.K_8
	Key9          = sdl.K_9
	KeyColon      = sdl.K_COLON
	KeySemicolon  = sdl.K_SEMICOLON
	KeyLess       = sdl.K_LESS
	KeyEquals     = sdl.K_EQUALS
	KeyGreater    = sdl.K_GREATER
	KeyQuestion   = sdl.K_QUESTION
	KeyAt         = sdl.K_AT

	KeyLeftBracket  = sdl.K_LEFTBRACKET
	KeyBackslash    = sdl.K_BACKSLASH
	KeyRightBracket = sdl.K_RIGHTBRACKET
	KeyCaret        = sdl.K_CARET
	KeyUnderscore   = sdl.K_UNDERSCORE
	KeyBackquote    = sdl.K_BACKQUOTE
	KeyA            = sdl.K_a
	KeyB            = sdl.K_b
	KeyC            = sdl.K_c
	KeyD            = sdl.K_d
	KeyE            = sdl.K_e
	KeyF            = sdl.K_f
	KeyG            = sdl.K_g
	KeyH            = sdl.K_h
	KeyI            = sdl.K_i
	KeyJ            = sdl.K_j
	KeyK            = sdl.K_k
	KeyL            = sdl.K_l
	KeyM            = sdl.K_m
	KeyN            = sdl.K_n
	KeyO            = sdl.K_o
	KeyP            = sdl.K_p
	KeyQ            = sdl.K_q
	KeyR            = sdl.K_r
	KeyS            = sdl.K_s
	KeyT            = sdl.K_t
	KeyU            = sdl.K_u
	KeyV            = sdl.K_v
	KeyW            = sdl.K_w
	KeyX            = sdl.K_x
	KeyY            = sdl.K_y
	KeyZ            = sdl.K_z

	KeyCapsLock = sdl.K_CAPSLOCK

	KeyF1  = sdl.K_F1
	KeyF2  = sdl.K_F2
	KeyF3  = sdl.K_F3
	KeyF4  = sdl.K_F4
	KeyF5  = sdl.K_F5
	KeyF6  = sdl.K_F6
	KeyF7  = sdl.K_F7
	KeyF8  = sdl.K_F8
	KeyF9  = sdl.K_F9
	KeyF10 = sdl.K_F10
	KeyF11 = sdl.K_F11
	KeyF12 = sdl.K_F12

	KeyPrintScreen = sdl.K_PRINTSCREEN
	KeyScrollLock  = sdl.K_SCROLLLOCK
	KeyPause       = sdl.K_PAUSE
	KeyInsert      = sdl.K_INSERT
	KeyHome        = sdl.K_HOME
	KeyPageUp      = sdl.K_PAGEUP
	KeyDelete      = sdl.K_DELETE
	KeyEnd         = sdl.K_END
	KeyPageDown    = sdl.K_PAGEDOWN
	KeyRight       = sdl.K_RIGHT
	KeyLeft        = sdl.K_LEFT
	KeyDown        = sdl.K_DOWN
	KeyUp          = sdl.K_UP

	KeyNumlockClear = sdl.K_NUMLOCKCLEAR
	KeyKpDivide     = sdl.K_KP_DIVIDE
	KeyKpMultiply   = sdl.K_KP_MULTIPLY
	KeyKpMinus      = sdl.K_KP_MINUS
	KeyKpPlus       = sdl.K_KP_PLUS
	KeyKpEnter      = sdl.K_KP_ENTER
	KeyKp1          = sdl.K_KP_1
	KeyKp2          = sdl.K_KP_2
	KeyKp3          = sdl.K_KP_3
	KeyKp4          = sdl.K_KP_4
	KeyKp5          = sdl.K_KP_5
	KeyKp6          = sdl.K_KP_6
	KeyKp7          = sdl.K_KP_7
	KeyKp8          = sdl.K_KP_8
	KeyKp9          = sdl.K_KP_9
	KeyKp0          = sdl.K_KP_0
	KeyKpPeriod     = sdl.K_KP_PERIOD

	KeyApplication   = sdl.K_APPLICATION
	KeyPower         = sdl.K_POWER
	KeyKpEquals      = sdl.K_KP_EQUALS
	KeyF13           = sdl.K_F13
	KeyF14           = sdl.K_F14
	KeyF15           = sdl.K_F15
	KeyF16           = sdl.K_F16
	KeyF17           = sdl.K_F17
	KeyF18           = sdl.K_F18
	KeyF19           = sdl.K_F19
	KeyF20           = sdl.K_F20
	KeyF21           = sdl.K_F21
	KeyF22           = sdl.K_F22
	KeyF23           = sdl.K_F23
	KeyF24           = sdl.K_F24
	KeyExecute       = sdl.K_EXECUTE
	KeyHelp          = sdl.K_HELP
	KeyMenu          = sdl.K_MENU
	KeySelect        = sdl.K_SELECT
	KeyStop          = sdl.K_STOP
	KeyAgain         = sdl.K_AGAIN
	KeyUndo          = sdl.K_UNDO
	KeyCut           = sdl.K_CUT
	KeyCopy          = sdl.K_COPY
	KeyPaste         = sdl.K_PASTE
	KeyFind          = sdl.K_FIND
	KeyMute          = sdl.K_MUTE
	KeyVolumeUp      = sdl.K_VOLUMEUP
	KeyVolumeDown    = sdl.K_VOLUMEDOWN
	KeyKpComma       = sdl.K_KP_COMMA
	KeyKpEqualsAS400 = sdl.K_KP_EQUALSAS400

	KeyAltErase   = sdl.K_ALTERASE
	KeySysReq     = sdl.K_SYSREQ
	KeyCancel     = sdl.K_CANCEL
	KeyClear      = sdl.K_CLEAR
	KeyPrior      = sdl.K_PRIOR
	KeyReturn2    = sdl.K_RETURN2
	KeySeparator  = sdl.K_SEPARATOR
	KeyOut        = sdl.K_OUT
	KeyOper       = sdl.K_OPER
	KeyClearAgain = sdl.K_CLEARAGAIN
	KeyCrSel      = sdl.K_CRSEL
	KeyExSel      = sdl.K_EXSEL

	KeyKP00               = sdl.K_KP_00
	KeyKP000              = sdl.K_KP_000
	KeyThousandsSeparator = sdl.K_THOUSANDSSEPARATOR
	KeyDecimalSeparator   = sdl.K_DECIMALSEPARATOR
	KeyCurrencyUnit       = sdl.K_CURRENCYUNIT
	KeyCurrencySubunit    = sdl.K_CURRENCYSUBUNIT
	KeyKPLeftParen        = sdl.K_KP_LEFTPAREN
	KeyKPRightParen       = sdl.K_KP_RIGHTPAREN
	KeyKPLeftBrace        = sdl.K_KP_LEFTBRACE
	KeyKPRightBrace       = sdl.K_KP_RIGHTBRACE
	KeyKPTab              = sdl.K_KP_TAB
	KeyKPBackspace        = sdl.K_KP_BACKSPACE
	KeyKPA                = sdl.K_KP_A
	KeyKPB                = sdl.K_KP_B
	KeyKPC                = sdl.K_KP_C
	KeyKPD                = sdl.K_KP_D
	KeyKPE                = sdl.K_KP_E
	KeyKPF                = sdl.K_KP_F
	KeyKPXor              = sdl.K_KP_XOR
	KeyKPPower            = sdl.K_KP_POWER
	KeyKPPercent          = sdl.K_KP_PERCENT
	KeyKPLess             = sdl.K_KP_LESS
	KeyKPGreater          = sdl.K_KP_GREATER
	KeyKPAmpersand        = sdl.K_KP_AMPERSAND
	KeyKPDblAmpersand     = sdl.K_KP_DBLAMPERSAND
	KeyKPVerticalBar      = sdl.K_KP_VERTICALBAR
	KeyKPDblverticalBar   = sdl.K_KP_DBLVERTICALBAR
	KeyKPColon            = sdl.K_KP_COLON
	KeyKPHash             = sdl.K_KP_HASH
	KeyKPSpace            = sdl.K_KP_SPACE
	KeyKPAt               = sdl.K_KP_AT
	KeyKPExclam           = sdl.K_KP_EXCLAM
	KeyKPMemStore         = sdl.K_KP_MEMSTORE
	KeyKPMemRecall        = sdl.K_KP_MEMRECALL
	KeyKPMemClear         = sdl.K_KP_MEMCLEAR
	KeyKPMemAdd           = sdl.K_KP_MEMADD
	KeyKPMemSubtract      = sdl.K_KP_MEMSUBTRACT
	KeyKPMemMultiply      = sdl.K_KP_MEMMULTIPLY
	KeyKPMemDivide        = sdl.K_KP_MEMDIVIDE
	KeyKPPlusMinus        = sdl.K_KP_PLUSMINUS
	KeyKPClear            = sdl.K_KP_CLEAR
	KeyKPClearEntry       = sdl.K_KP_CLEARENTRY
	KeyKPBinary           = sdl.K_KP_BINARY
	KeyKPOctal            = sdl.K_KP_OCTAL
	KeyKPDecimal          = sdl.K_KP_DECIMAL
	KeyKPHexadecimal      = sdl.K_KP_HEXADECIMAL

	KeyLCtrl  = sdl.K_LCTRL
	KeyLShift = sdl.K_LSHIFT
	KeyLAlt   = sdl.K_LALT
	KeyLGui   = sdl.K_LGUI
	KeyRCtrl  = sdl.K_RCTRL
	KeyRShift = sdl.K_RSHIFT
	KeyRAlt   = sdl.K_RALT
	KeyRGui   = sdl.K_RGUI

	KeyMode = sdl.K_MODE

	KeyAudioNext   = sdl.K_AUDIONEXT
	KeyAudioPrev   = sdl.K_AUDIOPREV
	KeyAudioStop   = sdl.K_AUDIOSTOP
	KeyAudioPlay   = sdl.K_AUDIOPLAY
	KeyAudioMute   = sdl.K_AUDIOMUTE
	KeyMediaSelect = sdl.K_MEDIASELECT
	KeyWWW         = sdl.K_WWW
	KeyMail        = sdl.K_MAIL
	KeyCalculator  = sdl.K_CALCULATOR
	KeyComputer    = sdl.K_COMPUTER
	KeyAcSearch    = sdl.K_AC_SEARCH
	KeyAcHome      = sdl.K_AC_HOME
	KeyAcBack      = sdl.K_AC_BACK
	KeyAcForward   = sdl.K_AC_FORWARD
	KeyAcStop      = sdl.K_AC_STOP
	KeyAcRefresh   = sdl.K_AC_REFRESH
	KeyAcBookmarks = sdl.K_AC_BOOKMARKS

	KeyBrightnessDown = sdl.K_BRIGHTNESSDOWN
	KeyBrightnessUp   = sdl.K_BRIGHTNESSUP
	KeyDisplaySwitch  = sdl.K_DISPLAYSWITCH
	KeyKbdillumToggle = sdl.K_KBDILLUMTOGGLE
	KeyKbdillumDown   = sdl.K_KBDILLUMDOWN
	KeyKbdillumUp     = sdl.K_KBDILLUMUP
	KeyEject          = sdl.K_EJECT
	KeySleep          = sdl.K_SLEEP
)
