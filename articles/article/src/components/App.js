import React, {useState} from "react";
import axios from "axios";

function CreateArticleForm(){
    const [title, setTitle] = useState("");
    const [body, setBody] = useState("");
    const [author, setAuthor] = useState("");
    const [category, setCategory] = useState("");

    const handleSubmit = async(e) =>{
        e.preventDefault();
        try{
            await axios.post("/api/articles",{
                title,
                body,
                author,
                category,
            });
            alert ("Article created successfully");
            setTitle("");
            setBody("");
            setAuthor("");
            setCategory("");
        }
        catch (err){
            console.error(err);
        }
    }
    return(
        <form onSubmit={handleSubmit}>
            <h2>Create new article</h2>
            <div>
                <label htmlFor="title">Title:</label>
                <input type="text" id="title" value={title} onChange={(e)=> setTitle(e.target.value)}/>

            </div>
            <div>
                <label htmlFor="body">Body:</label>
                <textarea id="body" value={body} onChange={(e) =>setBody(e.target.value)}/>

      </div>
      <div>
        <label htmlFor ="author">Author:</label>
        <input type="text" id="author" value={author} onChange={(e)=>setAuthor(e.target.value)}/>

      </div>
      <div>
        <label htmlFor ="category">Category:</label>
        <input type="text" id="category" value={category} onChange={(e)=>setCategory(e.target.value)}/>

      </div>
      <button type ="submit">Create</button>
        </form>
    );
}
function SearchArticlesForm() {
    const [searchQuery, setSearchQuery] =useState("");
    const [searchResults, setSearchResults] = useState({});

    const handleSubmit =async(e)=>{
        e.preventDefault();
        try{
            const res = await axios.get(`api/articles?search=${searchQuery}`);
            setSearchResults(res.data);

        }
        catch(err){
            console.error(err);
        }
    };
    return(
        <form onSubmit ={handleSubmit}>
        <h2>Search for articles</h2>
        <div>
            <label htmlFor="searchQuery">Search query:</label>
            <input type="text" id="searchQuery" value={searchQuery} onChange={(e)=> setSearchQuery (e.target.value)}/>

        </div>
        <button type="submit">Search</button>
        <div>
            {Object.keys(searchResults).map((category)=>(
                <div key ={category}>
                    <h3>{category}</h3>
                    <ul>
                        {searchResults[category].map((article)=>(
                            <li key ={article.ID}>
                                <h4>{article.Title}</h4>
                                <p>{article.Body}</p>
                                <p>Author:{article.Author}</p>
                                <p>Category: {article.Category}</p>

                            </li>
                        ))}
                    </ul>
                </div>
                
            ))}
        </div>


        </form>
    );
}
function App(){
    return(
        <div>
            <CreateArticleForm/>
            <SearchArticlesForm/>
        </div>
    )
}
export default App;