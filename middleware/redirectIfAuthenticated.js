export default function({ redirect }) {
  if (localStorage.getItem('loggedIn') === 'true') {
    return redirect('/dashboard')
  }
}
