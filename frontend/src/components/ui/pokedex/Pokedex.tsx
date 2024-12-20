import React, {PropsWithChildren} from "react";
import {LeftScreenContainer, RightScreenContainer} from "./PokedexScreens.tsx";
import {TabButton, TabButtonContainer} from "./TabButton.tsx";
import {KeyboardButton, KeyboardButtonContainer} from "./KeyboardButton.tsx";


export const PokedexRoot: React.FC<PropsWithChildren> = ({children}) => (
    <div
        className={"relative grid grid-cols-1 gap-0.5 sm:gap-2 md:grid-cols-2 md:grid-rows-7 md:gap-x-3 lg:gap-x-5 bg-red-700 px-3 sm:px-12 md:px-6 lg:px-8 rounded-xl max-w-md md:max-w-4xl lg:max-w-5xl xl:max-w-6xl 2xl:max-w-screen-2xl mx-1 xs:mx-auto pt-12 sm:pt-8 pb-6 sm:min-h-[726px] max-h-screen shadow-md dark:shadow-gray-800"}>
        {children}
        <div
            className="h-0.5 sm:h-1 rounded-full bg-slate-700 w-full md:col-span-2 mx-auto self-end mt-1 sm:mt-1.5"></div>
    </div>
);


const Pokedex = PokedexRoot as typeof PokedexRoot & {
    RightScreen: typeof RightScreenContainer;
    LeftScreen: typeof LeftScreenContainer;
    KeyboardButtons: typeof KeyboardButtonContainer;
    KeyboardButton: typeof KeyboardButton
    TabButtons: typeof TabButtonContainer;
    TabButton: typeof TabButton
}

Pokedex.RightScreen = RightScreenContainer
Pokedex.LeftScreen = LeftScreenContainer;
Pokedex.KeyboardButtons = KeyboardButtonContainer;
Pokedex.KeyboardButton = KeyboardButton
Pokedex.TabButtons = TabButtonContainer;
Pokedex.TabButton = TabButton;


export default Pokedex;



