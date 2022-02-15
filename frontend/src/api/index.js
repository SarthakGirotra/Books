import axios from "axios"

const API_User = axios.create({ baseURL: 'http://localhost:1323' })
const API_Books = axios.create({ baseURL: 'http://localhost:1322' })

export const login = (formdata) => API_User.post("/login", formdata).catch((error) => error.response)
export const signup = (formdata) => API_User.post("/signup", formdata).catch((error) => error.response)
export const userAPIHealth = () => API_User.get("/health").catch((error) => error.response)
export const getTopBooks = () => API_Books.get("/topBooks").catch((error) => error.response)
export const likeStory = (ids) => API_Books.post("/Like", ids)