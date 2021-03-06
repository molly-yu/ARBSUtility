import {FETCH_CAMERAS, NEW_CAMERA, DELETE_CAMERA} from '../actions/types';

const initialState = {
    items:[],
    item:{}
};

export default function(state = initialState, action){
    switch(action.type){
        case FETCH_CAMERAS:
            return {
                ...state,
                 items: action.payload
            };
        case NEW_CAMERA:
            return {
                ...state,
                 item: action.payload
            };
        case DELETE_CAMERA:
        return {
            ...state,
            items: state.items.filter(item1 => item1.id !== action.payload)

            };
        default:
            return state;
    }
}