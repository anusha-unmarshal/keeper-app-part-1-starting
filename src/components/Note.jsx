import React from "react";
// import EdiText from "react-editext";
function Note(props) {
    function handleClick() {
        props.onDelete(props.id);
    }
    function editNote() {
        props.onEdit(props.id);
    }
    return (
        <div className="note">
            <h1>{props.title}</h1>
            <p>{props.content}</p>
            {/* <EdiText 
                type="text"
                value = {props.content}
                onSave = {onSave} */}
            {/* /> */}
            <button onClick={editNote}>EDIT</button>
            <button onClick={handleClick}>DELETE</button>
        </div>
    );
}
export default Note;