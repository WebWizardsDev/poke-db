import React from "react";
import {MoveType} from '../../models/pokemon.ts'
import {TypeListItem} from "./TypeListItem.tsx";


interface TypeListProps {
    moveTypes?: MoveType[];
    listIndex: number;
    onTypeSelect: (index: number) => void
}


export const TypeList: React.FC<TypeListProps> = ({moveTypes, listIndex, onTypeSelect}) => {


    return (<div
        className={`grid grid-cols-2 items-stretch justify-center gap-1 xs:gap-2.5 sm:gap-3 md:gap-4 sm:my-3`}>
        {moveTypes?.map((moveType, index) => {
            return <TypeListItem active={index === listIndex} index={index} type={moveType} key={`type_${index}`}
                                 onTypeSelect={onTypeSelect}/>
        })}
    </div>)
}