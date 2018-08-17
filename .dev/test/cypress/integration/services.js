context('Service Tests', () => {

    beforeEach(function () {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should view services', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('tr').should('have.length', 6)
        cy.title().should('eq', 'Statup | Services')
    })

    it('should create services', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('input[name="name"]').type('Google.com')
        cy.get('input[name="domain"]').type('https://google.com')
        cy.get('input[name="interval"]').type('30')
        cy.get('form').submit()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 7)
    })

    it('should delete a service', () => {
        cy.visit('http://localhost:8080/services')
        cy.get(':nth-child(5) > .text-right > .btn-group > .btn-danger').click()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 6)
    })


});