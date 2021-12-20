import argparse as ap
import numpy as np
from pynput.keyboard import Controller as KBController
from pynput.mouse import Button as MButton, Controller as MController
import random
import time
from typing import Tuple


def progress_bar(progress: int, total: int, prefix: str = "", suffix: str = "", 
                 decimals: int = 1, bar_length: int = 30, bar_char: str = "█") -> None:
    """Create and print a progress bar.

    Args:
        progress (int): the progress toward the goal
        total (int): the goal
        prefix (str): the string to be printed before the progress bar
        suffix (str): the string to be printed after the progress bar
        decimals (int): the number of decimal places to display in the 
            percentage progress display
        bar_length (int): the horizontal width of the progress bar
        bar_char (str): the character to use for units on the progress bar
    """

    percent = progress / total
    bar_count = int(percent * bar_length)
    bars = bar_char * bar_count
    spaces = " " * (bar_length - bar_count)

    print(f"{prefix}|{bars}{spaces}| {percent:.{decimals}%}{suffix}")

    return None

def estimate_time(craft_count: int, craft_progress: int, click_delay: float, 
                  craft_delay: float, craft_length: float) -> Tuple[int, int]:
    """Estimate the time remaining in minutes and seconds.

    Args:
        craft_count (int): the number of crafts for the session
        craft_progress (int): the number of crafts completed
        click_delay (float): the delay after clicking
        craft_delay (float): the delay after finishing a craft

    Returns:
        A tuple of the form (minutes, seconds) representing the remaining 
            time to complete scheduled crafts.
    """

    time_per_craft = (
        (0.001 + 0.05) + (0.001 + 1.2 * click_delay) + (0.001 + 0.1) 
         + (craft_length + 1.0) + (1.4 * craft_delay))
    
    crafts_remaining = craft_count - craft_progress

    time_remaining = crafts_remaining * time_per_craft

    minutes = int(time_remaining / 60.0)
    seconds = int(time_remaining - minutes * 60.0)

    return (minutes, seconds)


if __name__ == "__main__":

    # setup command-line argument parsing
    parser = ap.ArgumentParser(
        description="Automate basic crafting task with macro. For crafts "
                    "spanning multiple macros, I recommend using the "
                    "/nextmacro command provided by the Macro Chain Dalamud "
                    "plugin.")
    parser.add_argument("craft_count", type=int, 
                        help="number of items to craft")
    parser.add_argument("craft_length", type=float, 
                        help="the duration of the craft macro")
    parser.add_argument("macro_key", type=str, 
                        help="the key corresponding to the crafting macro")
    parser.add_argument("--click_delay", "-clkdel", type=float, default=1.0,
                        help="the time (in seconds) to wait after automated clicking")
    parser.add_argument("--craft_delay", "-crftdel", type=float, default=4.0,
                        help="the time (in seconds) to wait after finishing an automated crafting macro")
    parser.add_argument("--initial_delay", "-initdel", type=float, default=10.0,
                        help="the time (in seconds) before starting the craft automation")
    args = parser.parse_args()
    
    # setup random number generation
    rng = np.random.default_rng(random.randint(0,100000000))

    # setup input management
    mouse = MController()
    kb = KBController()

    # wait for initial delay
    time.sleep(args.initial_delay)

    # start automated crafting loop
    items_crafted = 0
    while items_crafted < args.craft_count:
        # estimate and display remaining craft time
        estimated_time_remaining = estimate_time(args.craft_count, items_crafted, args.click_delay, args.craft_delay, args.craft_length)
        print(f"Estimated time left for craft completion: {estimated_time_remaining[0]}m{estimated_time_remaining[1]}s.")

        # click the button to start crafting
        mouse.press(MButton.left)
        time.sleep(0.001 + rng.random() * 0.1)
        mouse.release(MButton.left)

        # wait before starting the crafting macro
        time.sleep(args.click_delay + rng.random() * 0.4 * args.click_delay)

        # start the crafting macro
        kb.press(args.macro_key)
        time.sleep(0.001 + rng.random() * 0.2)
        kb.release(args.macro_key)

        # wait until the crafting macro is finished
        time.sleep(args.craft_length + rng.random() * 2.0)

        # increment crafted item count and display progress bar
        items_crafted += 1
        progress_bar(
            items_crafted, args.craft_count, prefix="Progress: ", 
            suffix=f" {items_crafted}/{args.craft_count} Items Crafted")

        # wait before starting crafting for next item
        time.sleep(args.craft_delay + rng.random() * 0.4 * args.craft_delay)
    
    print("Crafting complete!")

    # wait for the user to press a key to quit.
    input("Press any key to quit.")
