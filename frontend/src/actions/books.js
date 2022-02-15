import * as api from "../api"

export const getTopBooks = () => async (dispatch) => {
    try {
        const { data } = await api.getTopBooks()
        dispatch({ type: "TOP", data })
    } catch (error) {
        console.log(error)
    }

}

export const likeBook = (ids) => async (dispatch) => {
    try {
        dispatch({ type: "LIKE", ids })
        await api.likeStory(ids)

    } catch (error) {
        console.log(error)
    }
}

