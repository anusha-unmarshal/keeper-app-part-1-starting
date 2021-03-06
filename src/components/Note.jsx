import React, { useState } from "react";
// import EdiText from "react-editext";
function Note(props) {

    const [edit, setEditValue] = useState(false);
    function handleClick() {
        props.onDelete(props.id);
    }
    function editNote() {
        setEditValue(true);
    }
    const [content, updateContent] = useState(props.content);
    function handleChange(event) {
        updateContent(event.target.value);
        // props.onSave(props.id,content);
    }

    function saveNote(){
        props.onSave(props.id,content);
        setEditValue(false);

    }
    return (
        <div className="note">
            {/* <p>{props.id}</p> */}
            <h1>{props.title}</h1>
            {!edit ?
            <p>{props.content}</p> :
            <input onChange={handleChange} type="text" value={content} name="editedContent" contentEditable/>
            }
            {/* <EdiText 
                type="text"
                value = {props.content}
                onSave = {onSave} */}
            {/* /> */}
            {!edit ?
                <button onClick={editNote}>EDIT</button> :
                <button onClick={saveNote}>SAVE</button>
            }
            
            <button onClick={handleClick}>DELETE</button>
        </div>
    );
}
export default Note;