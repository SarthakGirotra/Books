import * as api from "../api"


export const login = (formData) => async (dispatch) => {
    try {
        const data = await api.login(formData)
        if (data.status === 200) dispatch({ type: 'AUTH', data: data.data })
        else {
            dispatch({ type: "ERROR", data: data.data })
        }

    } catch (error) {
        console.log(error)
    }
}
export const signup = (formData) => async (dispatch) => {
    try {
        const data = await api.signup(formData)
        if (data.status === 200) dispatch({ type: 'AUTH', data: data.data })
        else {
            dispatch({ type: "ERROR", data: data.data })
        }


    } catch (error) {
        console.log(error)
    }
}
export const logout = () => (dispatch) => {
    dispatch({ type: 'LOGOUT' })

}

export const home = () => async (dispatch) => {
    const id = localStorage?.getItem('id')
    if (id) {
        try {
            dispatch({ type: 'HOME', payload: id })
        } catch (error) {
            console.log(error)
        }

    }

}