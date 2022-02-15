import React, { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { useDispatch, useSelector } from 'react-redux';
import { login, signup } from "../../actions/user"
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import InputAdornment from '@material-ui/core/InputAdornment';
import IconButton from '@material-ui/core/IconButton';
import Alert from '@mui/material/Alert';

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(18),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(3),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
    invalidEmail: {
        error: true,
        id: "standard-error-helper-text",
        label: "Error",
        helperText: "Invalid email"

    },
}));

export default function SignUp() {
    const err = useSelector((state) => state.error)
    const classes = useStyles();
    const [loggedin, setLogin] = useState(true)
    const dispatch = useDispatch();
    const [formData, setFormData] = useState({ email: "", password: "" })
    const [visible, setVisible] = useState(false)
    const [formV, setFormV] = useState({ email: false, password: false })
    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };
    const remError = () => {
        dispatch({ type: "REMOVE_ERROR" })
    }

    const validateEmail = (email) => {
        return String(email)
            .toLowerCase()
            .match(
                /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            );
    };
    const validatePass = (pass) => {
        return String(pass).match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{10,}$/)
    }
    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Avatar className={classes.avatar}>
                    <LockOutlinedIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                    {loggedin ? "Sign in" : "Sign up"}
                </Typography>
                <form className={classes.form} noValidate>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                variant="outlined"

                                //className={classes.invalidEmail}
                                error={formV.email}
                                helperText={formV.email ? "invalid email" : ""}
                                required
                                fullWidth
                                id="email"
                                label="Email"
                                name="email"
                                autoComplete="email"
                                value={formData.email}
                                onChange={(e) => {
                                    setFormData({ ...formData, email: e.target.value })
                                    if (formV.email) setFormV({ ...formV, email: validateEmail(formData.email) == null })

                                }}

                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                variant="outlined"
                                required
                                fullWidth
                                error={formV.password}
                                helperText={formV.password ? "min 10 characters with at least 1 lower case ,1 upper case, 1 number and 1 special character" : ""}
                                name="password"
                                label="Password"
                                type={visible ? "text" : "password"}
                                id="password"
                                value={formData.password}
                                autoComplete="current-password"
                                onChange={(e) => {

                                    setFormData({ ...formData, password: e.target.value })
                                    if (formV.password) setFormV({ ...formV, password: validatePass(formData.password) == null })
                                }}
                                onBlur={() => setFormV({ ...formV, password: validatePass(formData.password) == null })}
                                InputProps={{
                                    endAdornment:
                                        <InputAdornment position="end">
                                            <IconButton
                                                aria-label="toggle password visibility"
                                                onClick={() => setVisible(!visible)}
                                                onMouseDown={handleMouseDownPassword}
                                            >
                                                {visible ? <Visibility /> : <VisibilityOff />}
                                            </IconButton>
                                        </InputAdornment>
                                }}
                                onFocus={() => { setFormV({ ...formV, email: validateEmail(formData.email) == null }) }}
                            />
                        </Grid>
                    </Grid>
                    <Button
                        disabled={formData.email.length < 1 || formData.password.length < 1 || formV.email || validatePass(formData.password) == null}
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                        onClick={(e) => {
                            e.preventDefault();
                            setFormV({ ...formV, email: validateEmail(formData.email) == null })
                            setFormV({ ...formV, password: validatePass(formData.password) == null })
                            if (!formV.email && !formV.password) {
                                loggedin ? dispatch(login(formData)) : dispatch(signup(formData))
                                if (err) remError()
                            }

                        }}
                    >
                        {loggedin ? "Sign in" : "Sign up"}
                    </Button>
                    <Grid container justifyContent="flex-end">
                        <Grid item>
                            <Link href="#" variant="body2" onClick={(e) => { setLogin(!loggedin) }}>
                                {loggedin ? "Don't have an account? Sign up" : "Already have an account? Sign in"}
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
            {err != null ? (<Alert severity="error" sx={{ marginTop: "20px" }}>{err}</Alert>) : (null)}
        </Container >)




}