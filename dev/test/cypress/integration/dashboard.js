context('Dashboard Tests', () => {

    beforeEach(function() {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should view logs', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get(':nth-child(5) > .nav-link').click()
        cy.wait(10000)
        cy.get('#live_logs').should('contain', 'Service')
    })

    it('should view help', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get(':nth-child(6) > .nav-link').click()
        cy.title().should('eq', 'Statup | Help')
        cy.get('.col-12 > :nth-child(1)').should('contain', 'Statup')
    })

});