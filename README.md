# ffxivautocraft
Automate long, tedious crafts and get real-time progress updates in the terminal.

## Usage
```
usage: autocraft.exe [-h] [-click_delay CLICK_DELAY] [-craft_delay CRAFT_DELAY] [-initial_delay INITIAL_DELAY] craft_count craft_length macro_key

Automate basic crafting task with macro. For crafts spanning multiple macros, 
I recommend using the /nextmacro command provided by the Macro Chain Dalamud plugin.

positional arguments:
  craft_count           
        number of items to craft
  craft_length
        the duration of the craft macro
  macro_key             
        the key corresponding to the crafting macro

optional arguments:
  -h            
        show this help message and exit
  --click_delay CLICK_DELAY
        the time (in seconds) to wait after automated clicking
  --craft_delay CRAFT_DELAY
        the time (in seconds) to wait after finishing an automated crafting macro
  --initial_delay INITIAL_DELAY
        the time (in seconds) before starting the craft automation
```

## Example
Let's say you have your crafting macro (or starting macro for chained macros) placed on your hotbars bound to the 5 key, you are crafting 30 of the item, and you know each craft takes 20 seconds. To use autocraft for this situation, follow the steps below:

1. Launch FFXIV if you haven't already done so. 
2. Position yourself somewhere private, like an inn or FC room.
3. Open your crafting log and select the item you wish to craft.
5. Open a terminal (ideally on a second monitor where you can see the progress echoed in real-time) in the directory where you've placed autocraft.exe, and run the following command:
   `autocraft.exe 30 20 5`.
6. Alt+Tab back to FFXIV and hover your mouse over the center region of the `Synthesize` button.
7. Profit.

## Notes
- The easiest ways to determine the craft length are to time it with a stopwatch or to make the macro in teamcraft and use the reported craft duration.
- The crafts may be unable to continue if you run out of your starting material quality (HQ or NQ) and have to change the material quality in use for the craft.
