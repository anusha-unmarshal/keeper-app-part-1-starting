import React, { useState } from "react";
import Header from './Header';
import Note from "./Note";
import Footer from "./Footer";
import CreateArea from "./CreateArea";
import CenteredGrid from "./MaterialUI-example";


function App() {
    const [notes, setNotes] = useState([]);

    function addNote(newNote) {
        setNotes(prevNotes=> {
            return [...prevNotes, newNote];
        });
    }

    function deleteNote(id) {
        setNotes(prevNotes => {
            return prevNotes.filter((note, index) => {
                return index !== id; 
            });
        });
    }
    // function editNote(id){

    // }
    function onSaveText(id, editedContent) {
        // console.log(id);
        const newList = notes.map((item, index) => {
            if (index === id) {
                // console.log(item.id);
              const updatedItem = {
                ...item,
                content : editedContent
              };
       
              return updatedItem;
            }
       
            return item;
          });
       
          setNotes(newList);
    }
    
    return (
        <div>
        <Header />
        <CreateArea onAdd={addNote}/>
        {notes.map((note,index) => {
            return (<Note 
                key={index}
                id={index}
                title={note.title}
                content={note.content}
                onDelete={deleteNote}
                // onedit={editNote}
                onSave={onSaveText}
            />);
        })}
        <CenteredGrid />
        <Footer />
        </div>
    );
}

export default App;